package app

// BuildPipeline builds the HTTP pipeline
func (s *Server) Routes() {
	s.router.File("/", "web/build/index.html")
	s.router.Static("/static", "web/build/static")

	s.router.GET("/api/watch", s.Watch)
}
