package app

import (
	"net/http"
)

// BuildPipeline builds the HTTP pipeline
func (s *Server) Routes() {
	http.HandleFunc("/watch", s.Watch())
}
