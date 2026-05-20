package middleware

import (
	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/response"
	"github.com/ihorbryk/kalyna/router"
)

type Middleware func(h router.Handler) func(*request.Request) response.Response

func Apply(midls []Middleware, routes router.Routes) router.Routes {
	if len(midls) == 0 {
		return routes
	}

	mid := midls[0]
	midls = midls[1:]

	for _, route := range routes {
		route.Handler = mid(route.Handler)
	}

	return Apply(midls, routes)
}
