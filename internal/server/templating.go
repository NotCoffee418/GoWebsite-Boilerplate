package server

import (
	"embed"
	"html/template"
	"path/filepath"

	"github.com/NotCoffee418/GoWebsite-Boilerplate/internal/access"
	"github.com/NotCoffee418/GoWebsite-Boilerplate/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/NotCoffee418/GoWebsite-Boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
)

func initializeTemplates(engine *gin.Engine, templateFs embed.FS) {
	// Load HTML templates
	all_templates, err := utils.GetRecursiveFiles(
		templateFs,
		"templates",
		func(path string) bool { return filepath.Ext(path) == ".html" })
	if err != nil {
		log.Fatal("Error listing templates: ", err)
	}
	tmpl, err := template.ParseFiles(all_templates...)
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}

	// Set default template definitions (avoid crash on undefined template)
	setDefaultTemplateDefinitions(tmpl, config.DefaultTemplateDefinitions)

	// Register templates
	engine.SetHTMLTemplate(tmpl)
	access.Tmpl = tmpl
}

func setDefaultTemplateDefinitions(tmpl *template.Template, defs []config.DefaultTemplateDefinition) {
	for _, def := range defs {
		if tmpl.Lookup(def.Definition) == nil {
			_, err := tmpl.New(def.Definition).Parse(def.Content)
			if err != nil {
				log.Fatalf("Error parsing default template definition: %v", err)
			}
		}
	}
}
