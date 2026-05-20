package middleware

import (
	"log"
	"time"

	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/response"
	"github.com/ihorbryk/kalyna/router"
)

func RequestLogger(h router.Handler) func(r *request.Request) response.Response {
	return func(r *request.Request) response.Response {
		start := time.Now()

		response := h(r)

		end := time.Now()
		delta := end.Sub(start)

		log.Printf("%s %s %d %s", r.Method, r.URL.Path, response.StatusCode, delta)

		return response
	}
}
