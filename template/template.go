package template

import (
	"io/fs"
	"log"
	"path/filepath"
	"slices"
	"strings"
)

const BaseTemplateFolder = "templates"

type Context map[string]any

type TemplateName string
type TemplateFilePath string

type TemplateMap map[TemplateName]TemplateFilePath

func CollectTemplates(templateBaseDirs []string) TemplateMap {
	var files = TemplateMap{}

	for _, root := range templateBaseDirs {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err // propagate errors (e.g. permission denied)
			}

			if !d.IsDir() {
				parts := strings.Split(filepath.ToSlash(path), "/")
				i := slices.Index(parts, BaseTemplateFolder)

				if i != -1 {
					templateName := strings.Join(parts[i+1:], "/")
					files[TemplateName(templateName)] = TemplateFilePath(path)
				}
			}
			return nil
		})

		if err != nil {
			log.Printf("error to get templates in folder %s: %v", root, err)
		}
	}

	return files
}
