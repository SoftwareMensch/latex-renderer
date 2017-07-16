package config

import (
	"rkl.io/latex-renderer/di"
	"rkl.io/latex-renderer/services"
)

// Init initialise services
func Init() error {
	diContainer := di.GetContainer()

	// LaTeX rkl.io/latex-renderer
	diContainer.AddService("app.renderer", services.NewDocumentRenderer())

	return nil
}
