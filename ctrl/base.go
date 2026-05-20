package kalyna

import (
	"net/http"

	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/response"
)

type Base struct{}

func (ctrl Base) Get(r *request.Request) response.Response {
	return response.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
}

func (ctrl Base) Post(r *request.Request) response.Response {
	return response.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
}

func (ctrl Base) Put(r *request.Request) response.Response {
	return response.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
}

func (ctrl Base) Patch(r *request.Request) response.Response {
	return response.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
}

func (ctrl Base) Delete(r *request.Request) response.Response {
	return response.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
}

type Handlerer interface {
	Get(r *request.Request) response.Response
	Post(r *request.Request) response.Response
	Put(r *request.Request) response.Response
	Patch(r *request.Request) response.Response
	Delete(r *request.Request) response.Response
}

func AsHandler(h Handlerer) func(r *request.Request) response.Response {
	return func(r *request.Request) response.Response {
		switch r.Method {
		case http.MethodGet:
			return h.Get(r)
		case http.MethodPost:
			return h.Post(r)
		case http.MethodPut:
			return h.Put(r)
		case http.MethodPatch:
			return h.Patch(r)
		case http.MethodDelete:
			return h.Delete(r)
		default:
			return response.Response{
				StatusCode: http.StatusMethodNotAllowed,
			}
		}
	}
}
