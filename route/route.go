package route

import (
	"net/http"

	"github.com/carloct/slprofile/controller"
	hr "github.com/carloct/slprofile/route/middleware/httprouterwrapper"

	"github.com/julienschmidt/httprouter"
)

func LoadHTTP() http.Handler {
	return routes()
}

func routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", hr.HandlerFunc(controller.Index))

	return r
}

func middleware(h http.Handler) http.Handler {
	return h
}
