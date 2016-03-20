package auth

import (
	"net/http"

	"github.com/carloct/profile/shared/session"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := session.Instance(r)

		if sess.Values["id"] == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
