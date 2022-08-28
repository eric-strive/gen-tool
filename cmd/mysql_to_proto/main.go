package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/eric-strive/gen-tool/config"

	_ "github.com/go-sql-driver/mysql"
)

var typeArr = config.TypeArr

type Table struct {
	PackageModels string
	ServiceName   string
	Method        map[string]config.MethodDetail
	Comment       map[string]string
	Name          map[string][]Column
	Message       map[string]config.Detail
}

type Column struct {
	Field   string
	Type    string
	Comment string
}
type RpcServers struct {
	Models      string
	Name        string
	Funcs       []FuncParam
	MessageList []Message
}
type FuncParam struct {
	Name         string
	RequestName  string
	ResponseName string
	Method       string
	Description  string
	Path         string
	ProtoPath    string
}
type Message struct {
	Name          string
	MessageDetail []Field
}
type Field struct {
	TypeName     string
	AttrName     string
	FieldComment string
	Num          int
}

var tpl = config.ProtoTpl
var tablePre = config.ProtoTablePre
var filePath = config.ProtoFilePath
var dbName = config.ProtoDbName
var PackageModels = config.ProtoPackageModels
var ApiPrefix = config.ApiPrefix
var TableList = config.TableList
var ProtoPath = "recharge"

//var ServiceName = ""
func main() {
	file := filePath
	db, err := Connect(config.DriveName, fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.DbUser, config.DbPass, config.DbHost, config.DbPort)+dbName+"?charset=utf8mb4&parseTime=true")
	//Table names to be excluded
	exclude := map[string]int{"user_log": 1}
	if err != nil {
		fmt.Println(err)
	}
	if IsFile(file) {
		fmt.Fprintf(os.Stderr, "Fatal error: ", "file already exist")
		return
	}
	t := Table{}
	t.Message = config.ProtoMessage

	t.PackageModels = PackageModels
	t.ServiceName = StrFirstToUpper(PackageModels)
	t.Method = config.ProtoMethod
	t.TableColumn(db, dbName, exclude)
	t.Generate(file, tpl)
}

//Extract table field information
func (table *Table) TableColumn(db *sql.DB, dbName string, exclude map[string]int) {
	rows, err := db.Query("SELECT t.TABLE_NAME,t.TABLE_COMMENT,c.COLUMN_NAME,c.COLUMN_TYPE,c.COLUMN_COMMENT FROM information_schema.TABLES t,INFORMATION_SCHEMA.Columns c WHERE c.TABLE_NAME=t.TABLE_NAME AND t.`TABLE_SCHEMA`='" + dbName + "'")
	defer db.Close()
	defer rows.Close()
	var name, comment string
	var column Column
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: ", err)
		return
	}
	table.Comment = make(map[string]string)
	table.Name = make(map[string][]Column)
	for rows.Next() {
		rows.Scan(&name, &comment, &column.Field, &column.Type, &column.Comment)
		if _, ok := exclude[name]; ok {
			continue
		}
		if TableList != "" && !strings.Contains(TableList, name) {
			continue
		}
		if _, ok := table.Comment[name]; !ok {
			table.Comment[name] = comment
		}

		n := strings.Index(column.Type, "(")
		if n > 0 {
			column.Type = column.Type[0:n]
		}
		n = strings.Index(column.Type, " ")
		if n > 0 {
			column.Type = column.Type[0:n]
		}
		table.Name[name] = append(table.Name[name], column)
	}

	if err = rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: ", err)
		return
	}
	return
}

//Generate proto files in the current directory
func (table *Table) Generate(filepath, tpl string) {
	rpcservers := RpcServers{Models: table.PackageModels, Name: table.ServiceName}
	rpcservers.HandleFuncs(table)
	rpcservers.HandleMessage(table)
	tmpl, err := template.ParseFiles(tpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error1: ", err)
		return
	}
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0755)
	err = tmpl.Execute(file, rpcservers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error2: ", err)
		return
	}
}
func Connect(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(30)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	return db, err
}

func (r *RpcServers) HandleFuncs(t *Table) {
	var funcParam FuncParam
	for key, val := range t.Comment {
		k := StrFirstToUpper(strings.TrimPrefix(key, tablePre))
		for n, m := range t.Method {
			funcParam.Name = n + k
			funcParam.Method = m.Method
			funcParam.Description = fmt.Sprintf(m.Description, strings.Trim(val, "表"))
			funcParam.Path = ApiPrefix + strings.TrimPrefix(key, tablePre) + "/" + m.Path
			if m.Response.Cat == "existing" {
				funcParam.ResponseName = m.Response.Name
			} else {
				funcParam.ResponseName = k + m.Response.Name
			}
			funcParam.ProtoPath = ProtoPath
			funcParam.RequestName = k + m.Request.Name
			r.Funcs = append(r.Funcs, funcParam)
		}
	}
}

func (r *RpcServers) HandleMessage(t *Table) {
	message := Message{}
	field := Field{}
	var i int

	for key, value := range t.Name {
		k := StrFirstToUpper(strings.TrimPrefix(key, tablePre))

		for name, detail := range t.Message {
			message.Name = k + name
			message.MessageDetail = nil
			if detail.Cat == "all" {
				i = 1
				for _, f := range value {
					if isValueInList(f.Field, detail.FilterField) {
						/* 跳过此次循环 */
						continue
					}
					field.AttrName = f.Field
					field.TypeName = TypeMToP(f.Type)
					field.FieldComment = strings.Replace(f.Comment, "\n", "", -1) //去掉换行符
					field.Num = i
					message.MessageDetail = append(message.MessageDetail, field)
					i++
				}
			} else if detail.Cat == "custom" {
				i = 1
				for _, f := range detail.Attr {
					if f.TypeName == "repeated" {
						field.TypeName = "repeated " + k + "Request"
					} else {
						field.TypeName = f.TypeName
					}
					if f.ChildrenMessage != "" {
						field.TypeName = f.TypeName + " " + k + "" + config.ProtoMessage[f.ChildrenMessage].Name
					}
					field.AttrName = f.AttrName
					field.FieldComment = f.FieldComment
					field.Num = i
					message.MessageDetail = append(message.MessageDetail, field)
					i++
				}
			} else if detail.Cat == "one" {
				i = 1
				for _, f := range detail.Attr {
					field.TypeName = f.TypeName
					field.AttrName = f.AttrName
					field.FieldComment = f.FieldComment
					field.Num = i
					message.MessageDetail = append(message.MessageDetail, field)
					i++
				}
			}
			if len(detail.AppendField) > 0 {
				for _, va := range detail.AppendField {
					field.AttrName = va.AttrName
					field.TypeName = va.TypeName
					field.FieldComment = va.FieldComment
					field.Num = i
					message.MessageDetail = append(message.MessageDetail, field)
					i++
				}
			}
			if detail.Cat != "existing" {
				r.MessageList = append(r.MessageList, message)
			}
		}
	}
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func TypeMToP(m string) string {
	if _, ok := typeArr[m]; ok {
		return typeArr[m]
	}
	return "string"
}

func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for _, v := range temp {
		if len(v) > 0 {
			runes := []rune(v)
			upperStr += string(runes[0]-32) + string(runes[1:])
		}
	}
	return upperStr
}
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}
