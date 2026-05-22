package router

import (
	"fmt"
	"net/http"

	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/settings"
)

type Route struct {
	Path    string
	Name    string
	Handler Handler
}

func NewRoute(path, name string, handler Handler) *Route {
	return &Route{
		Path:    path,
		Name:    name,
		Handler: handler,
	}
}

type Routes []*Route

func New(routes ...*Route) Routes {
	return routes
}

func (r Routes) Append(routes ...*Route) Routes {
	return append(r, routes...)
}

func (r Routes) AddGroup(pathPrefix string, routes Routes) Routes {
	for _, route := range routes {
		route.Path = pathPrefix + route.Path
		r = append(r, route)
	}
	return r
}

func BindRoutesToServer(mux *http.ServeMux, routes Routes, s *settings.Settings) {
	for _, route := range routes {
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			request := request.New(r, s)
			response := route.Handler(request)

			w.WriteHeader(response.StatusCode)

			fmt.Fprintln(w, response.Body)
		})
	}
}
