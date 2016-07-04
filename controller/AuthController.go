package controller

import (
	"log"
	"net/http"
	"text/template"

	"github.com/cream-lab/vanilla/service"
)

func Signin(w http.ResponseWriter, req *http.Request) {
	var tmpl = template.Must(template.ParseGlob("views/login.html"))

	fbAuthConfig := config.GetFbAuthInfo()
	log.Println(fbAuthConfig.AppId)
	tmpl.ExecuteTemplate(w, "login.html", fbAuthConfig)
}
