package models

type TemplateData struct {
	StringData       map[string]string
	IntData          map[string]int
	floatData        map[string]float32
	CustomData       map[string]interface{}
	ValidationErrors map[string][]string
	flush            bool
	warning          string
	error            string
	success          string
	CSRFToken        string
}
