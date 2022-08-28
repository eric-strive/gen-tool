package config

const ProtoTpl = "./template/proto.go.tpl"
const ProtoTablePre = "ts_ad_" //操作去掉前缀
const ProtoFilePath = "/Users/wangwei/Work/Go/gen-tool/proto/db_ad_system.proto"
const ProtoDbName = "db_ad_system"
const ProtoPackageModels = "ad_system"
const ApiPrefix = "/internal/ad/cs/"

var TableList = "ts_ad_cs_advertising,ts_ad_cs_plan,ts_ad_cs_plan"

type MethodDetail struct {
	Request     Detail
	Response    Detail
	Method      string
	Path        string
	Description string
}

type AttrDetail struct {
	TypeName        string
	AttrName        string
	FieldComment    string
	ChildrenMessage string
}

type AppendField struct {
	TypeName     string
	AttrName     string
	FieldComment string
}
type Detail struct {
	Name        string
	Cat         string
	Attr        []AttrDetail
	FilterField []string
	AppendField []AppendField
}

//生成Proto文件信息配置
var ProtoMessage = map[string]Detail{
	"Filter": {
		Name:        "Filter",
		FilterField: []string{"id", "update_time", "create_by", "update_by"},
		AppendField: []AppendField{{"uint64", "num", "每页条数"},
			{"uint64", "offset", "偏移量"},
			{"string", "order_key", "排序key"},
			{"string", "order_desc", "排序要求"}},
		Cat: "all",
	},
	"GetInfo": {
		Name: "GetInfo",
		Cat:  "one",
		Attr: []AttrDetail{{
			TypeName:     "uint64",
			AttrName:     "id",
			FieldComment: "主键",
		}},
	},
	"Request": {
		Name: "Request",
		Cat:  "all",
	},
	"Create": {
		Name:        "Create",
		FilterField: []string{"id", "update_time", "create_time", "create_by", "update_by", "is_deleted"},
		AppendField: []AppendField{{"string", "operator", "操作人"}},
		Cat:         "all",
	},
	"Response": {
		Name: "common.Response",
		Cat:  "existing",
	},
	"ListResponse": {
		Name: "ListResponse",
		Cat:  "custom",
		Attr: []AttrDetail{
			{
				TypeName:     "uint32",
				AttrName:     "ret",
				FieldComment: "状态码",
			},
			{
				TypeName:     "string",
				AttrName:     "err",
				FieldComment: "异常信息",
			},
			{
				TypeName:        "repeated",
				AttrName:        "data",
				FieldComment:    "列表数据",
				ChildrenMessage: "DataInfo",
			},
		},
	},
	"DataInfo": {
		Name: "DataInfo",
		Cat:  "custom",
		Attr: []AttrDetail{
			{
				TypeName:     "repeated",
				AttrName:     "list",
				FieldComment: "查询数据列表",
			},
			{
				TypeName:     "int64",
				AttrName:     "total_num",
				FieldComment: "数据总条数",
			},
			{
				TypeName:     "int32",
				AttrName:     "num",
				FieldComment: "每页数据条数",
			},
			{
				TypeName:     "int32",
				AttrName:     "offset",
				FieldComment: "偏移量",
			},
		},
	},
}

//定义生成proto方法
var ProtoMethod = map[string]MethodDetail{
	"Get":    {Request: ProtoMessage["GetInfo"], Response: ProtoMessage["Request"], Method: "get", Path: "detail", Description: "获取%s详情"},
	"List":   {Request: ProtoMessage["Filter"], Response: ProtoMessage["ListResponse"], Method: "get", Path: "list", Description: "获取%s列表"},
	"Create": {Request: ProtoMessage["Create"], Response: ProtoMessage["Response"], Method: "post", Path: "add", Description: "新增%s"},
	"Update": {Request: ProtoMessage["Create"], Response: ProtoMessage["Response"], Method: "post", Path: "detail", Description: "更新%s"},
	"Delete": {Request: ProtoMessage["GetInfo"], Response: ProtoMessage["Response"], Method: "get", Path: "delete", Description: "删除%s"},
}
