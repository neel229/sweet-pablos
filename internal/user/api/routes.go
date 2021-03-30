package user

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// SetRoutes initializes all the routes the server will listen to
func (s *Server) SetRoutes() {
	s.r.Use(middleware.Logger)
	s.r.Route("/api/user", func(r chi.Router) {
		r.Get("/", s.GetUserByEmail())
		r.Post("/", s.CreateUser())
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetUser())
			r.Delete("/", s.DeleteUser())
			r.Put("/fname", s.UpdateFirstName())
			r.Put("/lname", s.UpdateLastName())
			r.Put("/email", s.UpdateEmail())
			r.Put("/password", s.UpdatePassword())
		})
	})
	s.r.Get("/api/users", s.ListUsers())
}
