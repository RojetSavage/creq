package req

import "net/http"

type RequestHandler struct {
	isReady bool
	http.Request
}

type ResponseHandler struct {
	*http.Response
}

type dismantledUrl struct {
	scheme   string
	user     string
	host     string
	port     string
	path     string
	query    string
	fragment string
}

type Config struct {
	DefaultUrl string `json:"defaultUrl"`
}
