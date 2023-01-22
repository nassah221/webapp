package helper

import (
	"bytes"
	"net/http"
	"text/template"
)

var tpl *template.Template

func LoadTemplates(pattern string) {
	tpl = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, pattern string, data any) {
	buf := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(buf, pattern, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
