package main

import (
	"fmt"
	"net/http"

	"github.com/brainattica/golang-jwt-authentication-api-sample/settings"
	"github.com/codegangsta/negroni"
	"github.com/kielan/go_account_svc/routers"
)

func main() {
	const port = "5000"
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
	fmt.Println("App ready. Listening at:", port)
}
