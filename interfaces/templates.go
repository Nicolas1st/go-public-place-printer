package interfaces

import "html/template"

type Templates interface {
	GetSignin() *template.Template
	GetSignup() *template.Template
}
