package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kielan/go_account_svc/controllers"
	"github.com/kielan/go_account_svc/core/authentication"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/check",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
