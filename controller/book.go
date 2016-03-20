package controller

import (
	"net/http"

	"github.com/carloct/profile/model"
	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"
)

func BookNew(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	v := view.New(r)
	v.Name = "book/new"

	sess.Save(r, w)

	v.Render(w)
}

func BookCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	_, err := model.BookCreate(sess.Values["id"].(int64), r.PostFormValue("title"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
