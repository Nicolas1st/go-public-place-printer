package pages

type Pages struct {
	Login      *LoginPage
	Signup     *SignupPage
	SubmitFile *SubmitFilePage
	Jobq       *JobqPage
}

// NewPages prebuilds all templates, for them to later be used only for
func NewPages(htmlTemplatesPath string) *Pages {
	return &Pages{
		Login:      NewLoginPage(htmlTemplatesPath),
		Signup:     NewSignupPage(htmlTemplatesPath),
		SubmitFile: NewSubmitFilePage(htmlTemplatesPath),
		Jobq:       NewJobqPage(htmlTemplatesPath),
	}
}
