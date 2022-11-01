package placekey

import "net/http"

const ()

// Key is the Placekey API key used globally in the binding
var Key string

type APIResponse struct {
	// Header contains a map of all HTTP header keys
	Header http.Header

	RawJSON []byte

	// QueryID contains a string that identifies the Placekey request if one was provided
	QueryID string

	// Status is a status code and message. e.g. "200 OK"
	Status string

	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}

type Backend interface {
	Call(method, path, key string) error
}
