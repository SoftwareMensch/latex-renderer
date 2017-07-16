package contract

// RendererInterface common rkl.io/latex-renderer interface
type RendererInterface interface {
	// Render render a dto
	Render(input interface{}) ([]byte, error)
}
