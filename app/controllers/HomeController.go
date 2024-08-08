package controllers

import (
	"forth/models"
	"forth/utils/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	AppConf models.AppConfig
}

func SetRepo(r *Repository) {
	Repo = r
}

func (m *Repository) HomeIndex(w http.ResponseWriter, r *http.Request) {
	var strings = make(map[string]string)
	strings["name"] = "Alireza"

	hasVisited := m.AppConf.Session.Get(r.Context(), "user") != ""
	m.AppConf.Session.Put(r.Context(), "user", r.RemoteAddr)
	var data = make(map[string]interface{})
	data["visited"] = hasVisited
	render.ViewRenderer(w, "home.tmpl", &models.TemplateData{
		StringData: strings,
		CustomData: data,
	}, r)
}
