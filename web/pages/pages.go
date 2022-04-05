package pages

import (
	"html/template"
)

type Pages struct {
}

type Templates struct {
	signup  *template.Template
	siginin *template.Template
}

func NewTemplates(htmlTemplatesPath string) *Templates {
	return &Templates{
		signup:  SignUpPage.BuildTemplate(htmlTemplatesPath),
		siginin: LoginPage.BuildTemplate(htmlTemplatesPath),
	}
}

// probably redundant, but the idea is to use
// dependency injection, to make
// the templates depent on views
func (t *Templates) GetSignup() *template.Template {
	return t.signup
}

func (t *Templates) GetLogin() *template.Template {
	return t.siginin
}