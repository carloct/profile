package controller

import (
	_ "fmt"
	"net/http"

	"github.com/carloct/profile/model"
	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"
)

// Displays the default home page
func Index(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	closets, err := model.Closets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Save(r, w)

	v := view.New(r)
	v.Name = "index"
	v.Vars["Closets"] = closets
	v.Render(w)
}

func CreateBook() {

}
