package auth

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	auth "github.com/neel229/sweet-pablos/internal/auth/service"
	"github.com/neel229/sweet-pablos/util"
)

// Server contains a router and a access token interface
type Server struct {
	r *chi.Mux
	c util.Config
	t auth.TokenInterface
}

func NewServer(config util.Config) (*Server, error) {
	tokenInterface, err := auth.NewJWT(config.TokenSymKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Server{
		r: chi.NewRouter(),
		c: config,
		t: tokenInterface,
	}, nil
}

func (s *Server) StartServer() {
	var err error
	s.c, err = util.LoadConfig("./config/auth")
	if err != nil {
		log.Fatalf("error loading config file: %v", err)
	}
	if err := http.ListenAndServe(":"+s.c.LisAddr, s.r); err != nil {
		log.Fatalf("error starting the server at port: %v", s.c.LisAddr)
	}
}
