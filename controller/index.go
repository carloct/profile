package controller

import (
	"net/http"

	"github.com/carloct/profile/shared/view"
)

// Displays the default home page
func Index(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "index"
	v.Render(w)
}
