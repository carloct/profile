package controller

import (
	"html/template"
	"net/http"
)

// Displays the default home page
func Index(w http.ResponseWriter, r *http.Request) {

	// Display the view
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, nil)
}
