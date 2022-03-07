package interfaces

import "html/template"

type Templates interface {
	GetLogin() *template.Template
	GetSignup() *template.Template
}
