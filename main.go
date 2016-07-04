package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/cream-lab/vanilla/route"
	"github.com/cream-lab/vanilla/service"
)

func main() {
	router := route.Router()

	n := negroni.Classic()
	n.UseHandler(router)

	// load app config file
	config.LoadAppConfig()

	http.ListenAndServe(":8080", n)
}
