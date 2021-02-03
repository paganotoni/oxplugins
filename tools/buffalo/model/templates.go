package model

var modelTemplate string = `package models

import (
	{{- range $i := .Imports }}
	"{{$i}}"
	{{- end }}
)

// {{ .Name.Proper.String }} model struct
type {{ .Name.Proper.String }} struct {
	{{- range $attr := .Attrs }}
	{{ $attr.Name.Pascalize }}	{{$attr.GoType }} ` + "`" + `json:"{{ $attr.Name.Underscore }}" db:"{{ $attr.Name.Underscore }}"` + "`" + `
	{{- end }}
}

// {{ .Name.Proper.Pluralize }} array model struct of {{ .Name.Proper.String }}
type {{ .Name.Proper.Pluralize }} []{{ .Name.Proper.String }}

// String converts the struct into a string value
func ({{ .Char }} {{ .Name.Proper.String }}) String() string {
	return fmt.Sprintf("%+v\n", {{ .Char }})
}
`

var modelTestTemplate string = `package models

func (ms *ModelSuite) Test_{{ .Name.Proper.String }}() {
	ms.Fail("This test needs to be implemented!")
}`
