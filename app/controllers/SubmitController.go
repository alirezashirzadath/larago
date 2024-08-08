package controllers

import (
	"forth/models"
	"forth/utils/render"
	"forth/utils/validation"
	"net/http"
)

func (m *Repository) GetSubmitController(w http.ResponseWriter, r *http.Request) {

	render.ViewRenderer(w, "submit.tmpl", &models.TemplateData{}, r)

}

func (m *Repository) PostSubmitController(w http.ResponseWriter, r *http.Request) {
	var validationRules = make(map[string]string)
	validationRules["username"] = "required|min:5|max:15"
	validationRules["password"] = "required|min:8|max:20"
	validationErrors := validation.FormValidate(r, validationRules)
	if len(validationErrors) > 0 {
		render.ViewRenderer(w, "submit.tmpl", &models.TemplateData{
			ValidationErrors: validationErrors,
		}, r)
	}
}
