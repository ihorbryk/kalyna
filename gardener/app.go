package gardener

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateApp(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get current working directory: %v", err)
	}

	dirs := []string{
		filepath.Join(cwd, "apps", name, "templates"),
		filepath.Join(cwd, "apps", name, "ctrls"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create %s: %v", dir, err)
		}
	}

	moduleName := readModuleName(cwd)

	content := fmt.Sprintf("package %s\n\nimport (\n\t\"github.com/ihorbryk/kalyna/apps\"\n\t\"github.com/ihorbryk/kalyna/settings\"\n)\n\nfunc App() *apps.App {\n\tapp := apps.New()\n\tapp.SetSetup(func(stt *settings.Settings) {\n\t\tdir := app.AppTemplatesDir()\n\t\tstt.TemplatesDirs = append(stt.TemplatesDirs, dir)\n\t})\n\treturn app\n}\n", name)
	writeFile(filepath.Join(cwd, "apps", name, "app.go"), content)

	content = fmt.Sprintf("package %s\n\nimport (\n\t\"%s/apps/%s/ctrls\"\n\n\tctrl \"github.com/ihorbryk/kalyna/ctrl\"\n\t\"github.com/ihorbryk/kalyna/router\"\n)\n\nfunc Routes() router.Routes {\n\treturn router.New(\n\t\trouter.NewRoute(\"/index\", \"test-ctrl\", ctrl.AsHandler(ctrls.IndexCtrl{})),\n\t)\n}", name, moduleName, name)
	writeFile(filepath.Join(cwd, "apps", name, "routes.go"), content)

	content = "package ctrls\n\nimport (\n\tctrl \"github.com/ihorbryk/kalyna/ctrl\"\n\t\"github.com/ihorbryk/kalyna/request\"\n\t\"github.com/ihorbryk/kalyna/response\"\n\t\"github.com/ihorbryk/kalyna/template\"\n)\n\ntype IndexCtrl struct {\n\tctrl.Base\n}\n\nfunc (ctrl IndexCtrl) Get(r *request.Request) response.Response {\n\treturn template.Render(\n\t\tr,\n\t\t\"index.html\",\n\t\tmap[string]any{\"Title\": \"Hello world!!!\"},\n\t)\n}"
	writeFile(filepath.Join(cwd, "apps", name, "ctrls", "index.go"), content)

	content = fmt.Sprintf("<H1>Index page from app \"%s\"</H1>", name)
	writeFile(filepath.Join(cwd, "apps", name, "templates", "index.html"), content)
}

func readModuleName(cwd string) string {
	data, err := os.ReadFile(filepath.Join(cwd, "go.mod"))
	if err != nil {
		log.Fatalf("cannot read go.mod: %v", err)
	}
	for _, line := range strings.SplitN(string(data), "\n", 10) {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	log.Fatal("cannot find module name in go.mod")
	return ""
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatalf("failed to write %s: %v", path, err)
	}
}
