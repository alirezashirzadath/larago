package models

import (
	my_template "forth/models/template"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InProduction   bool
	TemplateConfig my_template.TemplateConfig
	Session        *scs.SessionManager
}
