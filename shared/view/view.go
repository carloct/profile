package view

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	viewInfo           View
	childTemplates     []string
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
)

type Template struct {
	Root string `json:"Root"`
}

type View struct {
	BaseURI   string
	Extension string
	Folder    string
	Name      string
	Layout    string
	Caching   bool
	Vars      map[string]interface{}
	request   *http.Request
}

func Configure(vi View) {
	viewInfo = vi
}

func LoadTemplates(rootTemp string) {
	rootTemplate = rootTemp
}

func New(req *http.Request) *View {

	v := &View{}

	v.BaseURI = viewInfo.BaseURI
	v.Extension = viewInfo.Extension
	v.Folder = viewInfo.Folder
	v.Layout = viewInfo.Layout
	v.Name = viewInfo.Name
	v.request = req

	return v
}

func (v *View) Render(w http.ResponseWriter) {
	tc, ok := templateCollection[v.Name]

	if !ok || !viewInfo.Caching {

		templateList := make([]string, 0)
		templateList = append(templateList, "layouts/"+v.Layout)
		templateList = append(templateList, v.Name)

		fmt.Printf("%+v", templateList)

		for i, name := range templateList {
			path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
			if err != nil {
				http.Error(w, "Template path error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			templateList[i] = path
		}

		templates, err := template.New(v.Name).ParseFiles(templateList...)
		if err != nil {
			http.Error(w, "Template parse error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		templateCollection[v.Name] = templates

		tc = templates

		err = tc.ExecuteTemplate(w, v.Layout+"."+v.Extension, v.Vars)
		if err != nil {
			http.Error(w, "Template file error: "+err.Error(), http.StatusInternalServerError)
		}

	}
}
