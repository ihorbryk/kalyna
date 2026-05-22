package pongo2

import (
	"maps"

	"github.com/flosch/pongo2/v6"
)

type Renderer struct{}

func (pr *Renderer) Render(filePath string, context map[string]any) (string, error) {
	tmpl, err := pongo2.FromCache(filePath)
	if err != nil {
		return "", err
	}

	pongonContext := pongo2.Context{}
	maps.Copy(pongonContext, context)

	return tmpl.Execute(pongonContext)
}
