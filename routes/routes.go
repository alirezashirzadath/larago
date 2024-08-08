package routes

import (
	"forth/app/controllers"
	"forth/app/middlewares"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middlewares.SessionMiddleware)
	mux.Use(middlewares.CSRFMiddleware)
	mux.Get("/", controllers.Repo.HomeIndex)
	mux.Get("/submit", controllers.Repo.GetSubmitController)
	mux.Post("/submit", controllers.Repo.PostSubmitController)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))
	return mux
}
