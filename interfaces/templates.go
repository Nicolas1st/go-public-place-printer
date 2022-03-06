package interfaces

import "html/template"

type Templates interface {
	GetSigninTemplate() *template.Template
	GetSignupTemplate() *template.Template
}
