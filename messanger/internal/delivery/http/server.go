package http

import (
	"net/http"
	"time"
)

type HTTPServer struct {
	Server *http.Server
}

func CreateHTTPServer(port string, handler http.Handler) *HTTPServer {
	server := &HTTPServer{
		Server: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    1 * time.Second,
			WriteTimeout:   1 * time.Second,
		},
	}
	return server
}

func (h *HTTPServer) Run() error {
	return h.Server.ListenAndServe()
}
