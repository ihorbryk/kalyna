package request

import (
	"net/http"

	"github.com/ihorbryk/kalyna/settings"
)

type Request struct {
	*http.Request
	Settings *settings.Settings
}

func New(r *http.Request, s *settings.Settings) *Request {
	return &Request{Request: r, Settings: s}
}
