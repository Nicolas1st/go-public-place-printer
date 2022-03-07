package templates

import (
	"html/template"
	"printer/templates/pages"
)

type Pages struct {
}

type Templates struct {
	signup  *template.Template
	siginin *template.Template
}

func NewTemplates() *Templates {
	return &Templates{
		signup:  pages.SignUpPage.BuildTemplate(),
		siginin: pages.SignInPage.BuildTemplate(),
	}
}

// probably redundant, but the idea is to use
// dependency injection, to make
// the templates depent on views
func (t *Templates) GetSignup() *template.Template {
	return t.signup
}

func (t *Templates) GetSignin() *template.Template {
	return t.siginin
}
