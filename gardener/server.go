package gardener

import (
	"errors"
	"log"
	"net/http"

	"github.com/ihorbryk/kalyna/router"
	"github.com/ihorbryk/kalyna/settings"
)

func StartServer(stt *settings.Settings, routes router.Routes) {
	mux := http.NewServeMux()

	router.BindRoutesToServer(mux, routes, stt)

	server := &http.Server{
		Addr:    stt.Addr,
		Handler: mux,
	}

	log.Printf("Server starting on %s", stt.Addr)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
