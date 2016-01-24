package controller

import (
	"net/http"

	"github.com/carloct/profile/shared/session"
	"github.com/carloct/profile/shared/view"
)

func UserNew(w http.ResponseWriter, r *http.Request) {
	_ = session.Instance(r)

	v := view.New(r)
	v.Name = "register"

	v.Render(w)
}
