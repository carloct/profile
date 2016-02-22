package route

import (
	"net/http"

	"github.com/carloct/profile/controller"
	hr "github.com/carloct/profile/route/middleware/httprouterwrapper"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// *****************************************************************************
// Routes
// *****************************************************************************

func routes() *httprouter.Router {
	r := httprouter.New()

	// Serve static files, no directory browsing
	r.GET("/static/*filepath", hr.Handler(alice.
		New().
		ThenFunc(controller.Static)))

	r.GET("/register", hr.Handler(alice.New().ThenFunc(controller.UserNew)))
	r.POST("/register", hr.Handler(alice.New().ThenFunc(controller.UserCreate)))

	r.GET("/login", hr.Handler(alice.New().ThenFunc(controller.LoginGET)))
	r.POST("/login", hr.Handler(alice.New().ThenFunc(controller.LoginPOST)))

	r.POST("/closet/create", hr.Handler(alice.New().ThenFunc(controller.ClosetCreate)))
	// Home page
	r.GET("/", hr.Handler(alice.
		New().
		ThenFunc(controller.Index)))

	return r
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Prevents CSRF and Double Submits
	/*cs := csrfbanana.New(h, session.Store, session.Name)
	cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
	cs.ClearAfterUsage(true)
	cs.ExcludeRegexPaths([]string{"/static(.*)"})
	csrfbanana.TokenLength = 32
	csrfbanana.TokenName = "token"
	csrfbanana.SingleToken = false
	h = cs

	// Log every request
	h = logrequest.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)*/

	return h
}
