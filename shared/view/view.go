package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
	mutex              sync.RWMutex
	viewInfo           View
)

// Template root and children
type Template struct {
	Root string `json:"Root"`
}

// View attributes
type View struct {
	BaseURI   string
	Extension string
	Folder    string
	Layout    string
	Name      string
	Caching   bool
	Vars      map[string]interface{}
	request   *http.Request
}

// Configure will set the view information
func Configure(vi View) {
	viewInfo = vi
}

// Read the configuration
func ReadConfig() View {
	return viewInfo
}

// PrependBaseURI prepends the base URI to the string
func (v *View) PrependBaseURI(s string) string {
	return v.BaseURI + s
}

// New returns a new view
func New(req *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})

	v.BaseURI = viewInfo.BaseURI
	v.Extension = viewInfo.Extension
	v.Folder = viewInfo.Folder
	v.Layout = viewInfo.Layout
	v.Name = viewInfo.Name

	// Make sure BaseURI is available in the templates
	v.Vars["BaseURI"] = v.BaseURI

	// This is required for the view to access the request
	v.request = req

	return v
}

// Render a template to the screen
func (v *View) Render(w http.ResponseWriter) {

	tc, ok := templateCollection[v.Name]

	// If the template collection is not cached or caching is disabled
	if !ok || !viewInfo.Caching {

		templatesDir, err := filepath.Abs(v.Folder)
		if err != nil {
			fmt.Printf("%v", err)
		}

		layout, err := filepath.Abs(
			v.Folder +
				string(os.PathSeparator) +
				"layouts" +
				string(os.PathSeparator) +
				v.Layout + "." + v.Extension)
		if err != nil {
			log.Fatal(err)
		}

		content, err := filepath.Abs(v.Folder + string(os.PathSeparator) + v.Name + "." + v.Extension)
		if err != nil {
			log.Fatal(err)
		}

		includes, err := filepath.Glob(templatesDir + "/includes/*.html")
		if err != nil {
			log.Fatal(err)
		}

		// List of template names
		templateList := make([]string, 0)
		templateList = append(templateList, layout)
		templateList = append(templateList, content)
		templateList = append(templateList, includes...)

		// Determine if there is an error in the template syntax
		templates, err := template.New(v.Name).ParseFiles(templateList...)

		if err != nil {
			http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		templateCollection[v.Name] = templates

		// Save the template collection
		tc = templates
	}
	//Display the content to the screen
	err := tc.ExecuteTemplate(w, v.Layout+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}
