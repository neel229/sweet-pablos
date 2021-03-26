package auth

import "github.com/go-chi/chi/middleware"

func (s *Server) SetRoutes() {
	s.r.Use(middleware.Logger)
	s.r.Post("/login", s.LoginUser())
}
