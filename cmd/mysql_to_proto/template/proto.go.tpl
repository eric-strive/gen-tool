syntax = "proto3";
import "google/api/annotations.proto";
import "google/api/common.proto";
package {{.Models}};
option go_package = "proto/{{.Models}}";

// The {{.Models}} service definition.
service {{.Name}} {
 {{range .Funcs }} //{{.Description}}
 rpc {{.Name}}({{.RequestName}}) returns ({{.ResponseName}}) {
    option (google.api.http) = {
      {{.Method}}: "{{.Path}}" {{if (eq .Method "post")}}
      body: "*"  {{end}}
    };
 }
{{ end }}
}
{{range .MessageList }}
message {{.Name}} {
{{range .MessageDetail }} // @inject_tag: json:"{{.AttrName}}" form:"{{.AttrName}}"
//{{.FieldComment}}
{{.TypeName}} {{.AttrName}}={{.Num}};
{{ end }}}
{{ end }}

