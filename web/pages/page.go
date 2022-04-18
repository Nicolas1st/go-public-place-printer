package pages

import (
	"fmt"
	"html/template"
	"path"
)

type pageInfo struct {
	mainTemplateName  string
	templateFileNames []string
}

func buildTemplate(htmlTemplatesPath string, p pageInfo) *template.Template {
	tmpl := template.New(p.mainTemplateName)

	// prepending before all file names
	allFilePaths := []string{}
	for _, name := range p.templateFileNames {
		allFilePaths = append(
			allFilePaths,
			path.Join(htmlTemplatesPath, name),
		)
	}

	// building the template
	template, err := tmpl.ParseFiles(allFilePaths...)
	if err != nil {
		// The project should not start if the templates can not be built
		fmt.Println("Make you sure the path to the html templates was correctly specified")
		fmt.Println("The path should be relative to the main executable")
		panic(err)
	}

	return template
}
