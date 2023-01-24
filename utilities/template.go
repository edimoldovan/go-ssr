package utilities

// currently used templates
func GetTemplates() []string {
	f := GetExecutablePath()

	// for now we declare them here upfront
	var templateFiles = []string{
		// layouts
		f + "/templates/layouts/home.html",
		// partials
		f + "/templates/partials/head.html",
		f + "/templates/partials/header.html",
		f + "/templates/partials/body.html",
		f + "/templates/partials/footer.html",
	}
	return templateFiles
}
