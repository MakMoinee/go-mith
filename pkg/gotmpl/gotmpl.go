package gotmpl

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type viewService struct {
	Template *template.Template
	Layout   string
}

type IView interface {
	RenderWithTemplate(w http.ResponseWriter, data interface{}) error
	Render(w http.ResponseWriter, data interface{}) error
}

func NewView(layoutDir string) IView {
	svc := viewService{}
	filess := layoutFiles(layoutDir)
	t, err := template.ParseFiles(filess...)
	if err != nil {
		panic(err)
	}
	svc.Template = t
	return &svc

}

func NewViewWithTemplate(layout string, layoutDir string, files ...string) IView {
	svc := viewService{}
	files = append(layoutFiles(layoutDir), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	svc.Layout = layout
	svc.Template = t
	return &svc

}
func (vs *viewService) RenderWithTemplate(w http.ResponseWriter, data interface{}) error {
	return vs.Template.ExecuteTemplate(w, vs.Layout, data)
}
func (vs *viewService) Render(w http.ResponseWriter, data interface{}) error {
	return vs.Template.Execute(w, data)
}

func layoutFiles(layoutDir string) []string {
	files, err := filepath.Glob(layoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
