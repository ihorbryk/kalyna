package response

type Response struct {
	StatusCode int
	Body       string
}

func Html(body string, StatusCode int) Response {
	return Response{
		Body:       body,
		StatusCode: StatusCode,
	}
}
