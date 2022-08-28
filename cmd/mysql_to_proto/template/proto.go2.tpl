syntax = "proto3";
import "google/api/annotations.proto";
import "google/api/field_behavior.proto";
import "google/api/client.proto";
package {{.Models}};
option go_package = "proto/{{.Models}}";

// The {{.Models}} service definition.
service {{.Name}} {
 {{range .Funcs }} //{{.Description}}
 rpc {{.Name}}({{.RequestName}}) returns ({{.ResponseName}}) {
    option (google.api.http) = {
      {{.Method}}: "{{.Path}}"
    };
 }
{{ end }}
}
{{range .MessageList }}
message {{.Name}} {
{{range .MessageDetail }} // @inject_tag: json:"{{.AttrName}}" form:"{{.AttrName}}"
{{.TypeName}} {{.AttrName}}={{.Num}}; //{{.FieldComment}}
{{ end }}
}
{{ end }}

