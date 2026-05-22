package kalyna

import (
	"log"
	"os"

	"github.com/ihorbryk/kalyna/apps"
	"github.com/ihorbryk/kalyna/gardener"
	"github.com/ihorbryk/kalyna/router"
	"github.com/ihorbryk/kalyna/settings"
	"github.com/ihorbryk/kalyna/template/pongo2"
)

const cmdRunServer = "runserver"
const cmdCreateApp = "createapp"
const cmdPrintRoutes = "printroutes"

func Run(appList []*apps.App, routes router.Routes, opts ...settings.Option) {
	stt := &settings.Settings{
		Addr:             ":8000",
		TemplateRenderer: &pongo2.Renderer{},
	}

	for _, opt := range opts {
		opt(stt)
	}

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("a command is required (e.g. runserver)")
	}

	switch args[0] {
	case cmdRunServer:
		for _, app := range appList {
			app.Setup(stt)
		}
		gardener.StartServer(stt, routes)
	case cmdCreateApp:
		if len(args) < 2 {
			log.Fatal("createapp requires an app name (e.g. createapp myapp)")
		}
		gardener.CreateApp(args[1])
	case cmdPrintRoutes:
		gardener.PrintRoutes(routes)
	default:
		log.Fatalf("unknown command: %s", args[0])
	}
}
