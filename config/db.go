package config

const (
	DriveName = "mysql"
	DbHost    = "192.168.28.75"
	DbPort    = "3306"
	DbUser    = "root"
	DbPass    = "abcd1234"
)

var TypeArr = map[string]string{
	"tinyint":    "int32",
	"smallint":   "int32",
	"mediumint":  "int32",
	"int":        "int32",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "string",
	"timestamp":  "string",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}
