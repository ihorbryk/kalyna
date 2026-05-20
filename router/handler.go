package router

import (
	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/response"
)

type Handler func(r *request.Request) response.Response
