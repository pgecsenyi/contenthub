package server

import (
	"net/http"
)

// LogRecord Used by the built-in HTTP server to construct an HTTP response.
type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {

	return r.ResponseWriter.Write(p)
}

// WriteHeader sends an HTTP response header with status code.
func (r *LogRecord) WriteHeader(status int) {

	r.status = status
	r.ResponseWriter.WriteHeader(status)
}
