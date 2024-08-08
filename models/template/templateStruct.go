package my_template

import "html/template"

type TemplateConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
