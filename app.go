package kalyna

import (
	"log"
	"net/http"

	"github.com/ihorbryk/kalyna/router"
	"github.com/ihorbryk/kalyna/settings"
	pongo2 "github.com/ihorbryk/kalyna/template/pongo2"
)

func Run(routes router.Routes, opts ...settings.Option) {
	stt := &settings.Settings{
		Addr:             ":8000",
		TemplateRenderer: &pongo2.Renderer{},
	}

	for _, opt := range opts {
		opt(stt)
	}

	mux := http.NewServeMux()

	router.BindRoutesToServer(mux, routes, stt)

	server := &http.Server{
		Addr:    stt.Addr,
		Handler: mux,
	}

	log.Println("Server starting on " + stt.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
