package controller

import (
	"net/http"

	"github.com/carloct/profile/model"
)

func ClosetCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("name")

	err := model.ClosetCreate(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", 302)
}
