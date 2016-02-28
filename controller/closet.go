package controller

import (
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/carloct/profile/model"
	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func ClosetCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	userId := sess.Values["id"].(uint32)
	name := r.PostFormValue("name")

	err := model.ClosetCreate(userId, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sess.Save(r, w)

	http.Redirect(w, r, "/", 302)
}

func ClosetIndex(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	params := context.Get(r, "params").(httprouter.Params)

	closet_id, _ := strconv.Atoi(params.ByName("id"))

	closet, err := model.ClosetById(closet_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	closets, err := model.Closets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	sess.Save(r, w)

	v := view.New(r)
	v.Name = "closet"
	v.Vars["Closets"] = closets
	v.Vars["Closet"] = closet
	v.Render(w)

}
