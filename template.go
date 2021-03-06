package main

import (
	"bytes"
	"text/template"
)

var errorsTemplate = `
{{ range .Errors }}

func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := reasonerrors.FromError(err)
	return e.Reason == {{.Name}}_{{.Value}}.String() && e.Code == {{.HTTPCode}}
}

func Error{{.CamelValue}}() *reasonerrors.Error {
	return reasonerrors.New({{.HTTPCode}}, int({{.Name}}_{{.Value}}), {{.Name}}_{{.Value}}.String(), "{{.Message}}")
}

func Error{{.CamelValue}}f(format string, args ...interface{}) *reasonerrors.Error {
	 return reasonerrors.New({{.HTTPCode}}, int({{.Name}}_{{.Value}}), {{.Name}}_{{.Value}}.String(), fmt.Sprintf(format, args...))
}

{{- end }}
`

type errorInfo struct {
	Name       string
	Value      string
	HTTPCode   int
	CamelValue string
	Message    string
}

type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
