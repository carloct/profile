package controller

import (
	"net/http"

	"github.com/carloct/slprofile/shared/view"
)

func Index(w http.ResponseWriter, r *http.Request) {

	v := view.New(r)
	v.Name = "index"
	v.Render(w)
}
