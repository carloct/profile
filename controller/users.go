package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/carloct/profile/model"
	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"
)

func UserNew(w http.ResponseWriter, r *http.Request) {
	_ = session.Instance(r)

	v := view.New(r)
	v.Name = "register"

	v.Render(w)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	r.ParseForm()
	firstName := r.PostFormValue("first_name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	_, err := model.UserCreate(firstName, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func LoginGET(w http.ResponseWriter, r *http.Request) {

	v := view.New(r)
	v.Name = "login"
	v.Render(w)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	email := r.FormValue("email")
	password := r.FormValue("password")

	result, err := model.UserByEmail(email)
	if err == sql.ErrNoRows {
		sess.Save(r, w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err != nil {
		fmt.Printf("%+v", result)
		log.Println(err)
	} else if strings.Compare(result.Password, password) == 0 {
		session.Empty(sess)
		sess.Values["id"] = result.Id
		sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
