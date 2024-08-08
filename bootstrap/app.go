package main

import (
	"fmt"
	"forth/app/controllers"
	"forth/app/middlewares"
	"forth/models"
	"forth/routes"
	"forth/utils/exception"
	"forth/utils/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const serverAddr = "8000"

var appConf models.AppConfig

func Shirzad() {
	log.Println()
}
func main() {

	Shirzad()

	appConf.TemplateConfig.TemplateCache = render.CreateTemplateCache()
	appConf.TemplateConfig.UseCache = false
	appConf.InProduction = false
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie = scs.SessionCookie{
		Persist:  true,
		Secure:   appConf.InProduction,
		SameSite: http.SameSiteLaxMode,
	}
	middlewares.SetMainSession(session)
	appConf.Session = session
	render.SetAppConfig(&appConf)
	controllers.SetRepo(&controllers.Repository{
		AppConf: appConf,
	})
	fmt.Println(fmt.Sprintf("Listening to port " + serverAddr))
	server := http.Server{
		Addr:    ":" + serverAddr,
		Handler: routes.Routes(),
	}
	ServingErr := server.ListenAndServe()

	if ServingErr != nil {
		exception.Exception("Error in serving", ServingErr.Error())
	}
}
