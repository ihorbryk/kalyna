package pongo2

import (
	"maps"

	"github.com/flosch/pongo2/v6"
)

type Renderer struct{}

func (pr *Renderer) Render(filePath string, context map[string]any) (string, error) {
	tmpl := pongo2.Must(pongo2.FromFile(filePath))

	pongonContext := pongo2.Context{}
	maps.Copy(pongonContext, context)

	return tmpl.Execute(pongonContext)
}
