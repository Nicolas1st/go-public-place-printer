package views

import "printer/config"

type pages struct {
	Login       *page
	Signup      *page
	SubmitFile  *page
	Jobq        *page
	UserManager *page
	Profile     *page
}

// newPages prebuilds all templates, for them to later be used only for
func newPages(htmlTemplatesPath string, endpoints config.Endpoints) *pages {
	templateToExecute := "layout"

	commonPublicTemplates := []string{"layout.html", "navbar.html", "public-links.html", "footer.html"}
	commonPrivatecTemplates := []string{"layout.html", "navbar.html", "private-links.html", "footer.html"}
	commonAdminTemplates := []string{"layout.html", "navbar.html", "admin-links.html", "footer.html"}

	return &pages{
		Login:       buildPage("LoginPage", endpoints, htmlTemplatesPath, templateToExecute, append(commonPublicTemplates, "login.html")...),
		Signup:      buildPage("SignUpPage", endpoints, htmlTemplatesPath, templateToExecute, append(commonPublicTemplates, "signup.html")...),
		SubmitFile:  buildPage("SubmitFilePage", endpoints, htmlTemplatesPath, templateToExecute, append(commonPrivatecTemplates, "submit-file.html")...),
		UserManager: buildPage("UserManagerPage", endpoints, htmlTemplatesPath, templateToExecute, append(commonAdminTemplates, "user-manager.html")...),
		Profile:     buildPage("UserManagerPage", endpoints, htmlTemplatesPath, templateToExecute, append(commonAdminTemplates, "user-manager.html")...),
	}
}