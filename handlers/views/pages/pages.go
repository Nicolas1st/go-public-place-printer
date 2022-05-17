package pages

type publicPages struct {
	Login  *Page
	SignUp *Page
}

type privatePages struct {
	Printer *Page
}

type adminPages struct {
	Printer     *Page
	UserManager *Page
}

type Pages struct {
	Public  *publicPages
	Private *privatePages
	Admin   *adminPages
}

// newPages prebuilds all templates, for them to later be used only for
func NewPages(htmlTemplatesPath string) *Pages {
	templateToExecute := "layout"

	commonPublicTemplates := []string{"layout.html", "navbar.html", "public-links.html", "footer.html"}
	public := &publicPages{
		Login:  buildPage("LoginPage", htmlTemplatesPath, templateToExecute, append(commonPublicTemplates, "login.html")...),
		SignUp: buildPage("SignUpPage", htmlTemplatesPath, templateToExecute, append(commonPublicTemplates, "signup.html")...),
	}

	commonPrivatecTemplates := []string{"layout.html", "navbar.html", "private-links.html", "footer.html"}
	private := &privatePages{
		Printer: buildPage("PrinterPage", htmlTemplatesPath, templateToExecute, append(commonPrivatecTemplates, "submit-file.html")...),
	}

	commonAdminTemplates := []string{"layout.html", "navbar.html", "admin-links.html", "footer.html"}
	admin := &adminPages{
		Printer:     buildPage("PrinterPage", htmlTemplatesPath, templateToExecute, append(commonAdminTemplates, "submit-file.html")...),
		UserManager: buildPage("UserManagerPage", htmlTemplatesPath, templateToExecute, append(commonAdminTemplates, "user-manager.html")...),
	}

	return &Pages{
		Public:  public,
		Private: private,
		Admin:   admin,
	}
}
