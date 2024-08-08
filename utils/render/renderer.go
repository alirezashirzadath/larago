package render

import (
	"bytes"
	"forth/models"
	"forth/utils/exception"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
	"path/filepath"
)

var appConf *models.AppConfig

func SetAppConfig(config *models.AppConfig) {
	appConf = config
}

func ViewRenderer(w http.ResponseWriter, tmpl string, data *models.TemplateData, r *http.Request) {

	var templateCache map[string]*template.Template
	if appConf.TemplateConfig.UseCache {
		templateCache = appConf.TemplateConfig.TemplateCache
	} else {
		templateCache = CreateTemplateCache()
	}
	templateSet, exists := templateCache[tmpl]
	if !exists {
		exception.Exception("Template Not Found", "Template : "+tmpl)
	}
	var buf = new(bytes.Buffer)
	data.CSRFToken = nosurf.Token(r)
	executeErr := templateSet.Execute(buf, data)
	if executeErr != nil {
		exception.Exception("Error in execution", "Template : "+tmpl)
	}
	_, writeToResErr := buf.WriteTo(w)
	if writeToResErr != nil {
		exception.Exception("Error in writing buffer", "Template : "+tmpl)
	}
}

func CreateTemplateCache() map[string]*template.Template {
	var myCache = make(map[string]*template.Template)

	pages, pagesErr := filepath.Glob("./views/pages/*.tmpl")
	if pagesErr != nil {
		exception.Exception("Get Pages", pagesErr.Error())
	}
	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, templateSetErr := template.ParseFiles(page)
		if templateSetErr != nil {
			exception.Exception("Template Set Creation", templateSetErr.Error())
		}
		layouts, layoutsErr := filepath.Glob("./views/layouts/*.tmpl")
		if layoutsErr != nil {
			exception.Exception("Get Layouts", layoutsErr.Error())
		}
		if len(layouts) > 0 {
			var parseLayoutsErr error
			templateSet, parseLayoutsErr = templateSet.ParseGlob("./views/layouts/*.tmpl")
			if parseLayoutsErr != nil {
				exception.Exception("Parse Layouts", parseLayoutsErr.Error())
			}
		}
		myCache[name] = templateSet

	}
	return myCache

}
