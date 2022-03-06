package interfaces

import "html/template"

type Page interface {
	GetMainFileName() string
	GetAllFileNames() []string
	BuildTemplate() *template.Template
}
