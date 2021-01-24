package pkg

import (
	"text/template"

	"github.com/iancoleman/strcase"
)

var funcMap = map[string]interface{}{
	"snake_case": strcase.ToSnake,
	"variable": func(s string) string {
		return strcase.ToLowerCamel(s) + "Elem"
	},
}

var IntoSchemaMapTemplate = template.Must(template.New("into_schema_map").Funcs(funcMap).Parse(`
{{ define "simple" }}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.Type{{ .Type }},

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}

{{ define "simple_list" }}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Schema{Type: schema.Type{{ .Elem }}},

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}

{{ define "simple_set" }}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.TypeSet,
		Elem: &schema.Schema{Type: schema.Type{{ .Elem }}},

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}

{{ define "map" }}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.TypeMap,
		Elem: &schema.Schema{Type: schema.Type{{ .Elem }}},

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}


{{ define "complex_list" }}
	{{ .Name | variable }} := {{ .ElemTypeName }}{}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Resource{Schema: {{ .Name | variable }}.IntoSchemaMap()},

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}


{{ define "complex_set" }}
	{{ .Name | variable }} := {{ .ElemTypeName }}{}
	m["{{ .Name | snake_case }}"] = &schema.Schema{
		Type: schema.TypeSet,
		Elem: &schema.Resource{Schema: {{ .Name | variable }}.IntoSchemaMap()},
		Set:  tfutils.GetHashFunc(&{{ .Name | variable }}),

		{{ with .Default }}
		Default: {{ . }},
		{{ end }}

		{{ if .Required }}
		Required: true,
		{{ else }}
		Optional: true,
		{{ end }}

		{{ with .Description }}
		Description: {{ . }},
		{{ end }}
	}
{{ end }}

func (s *{{ .StructName }}) IntoSchemaMap() map[string]*schema.Schema {
	m := make(map[string]*schema.Schema)

	{{ range .Fields }}

	{{ if (eq .Template "simple") }}{{ template "simple" . }}{{ end }}
	{{ if (eq .Template "simple_list") }}{{ template "simple_list" . }}{{ end }}
	{{ if (eq .Template "simple_set") }}{{ template "simple_set" . }}{{ end }}
	{{ if (eq .Template "map") }}{{ template "map" . }}{{ end }}
	{{ if (eq .Template "complex_list") }}{{ template "complex_list" . }}{{ end }}
	{{ if (eq .Template "complex_set") }}{{ template "complex_set" . }}{{ end }}

	{{ end }}

	return m
}

`))
