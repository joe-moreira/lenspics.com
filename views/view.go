package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// LayoutDir is cool
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

// View to be exported
type View struct {
	Template *template.Template
	Layout   string
}

// Render is cool
func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// NewView opa
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// layoutfiles returns a slice of strings : the files
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
