package user

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	db "github.com/neel229/sweet-pablos/internal/user/db/sqlc"
	"github.com/neel229/sweet-pablos/util"
)

// Server contains a router and a db conn
type Server struct {
	r     *chi.Mux
	store *db.DBStore
}

// NewServer creates a new instance of the server
func NewServer(s *db.DBStore) *Server {
	return &Server{
		r:     chi.NewRouter(),
		store: s,
	}
}

// StartServer starts a new server at the given port
func (s *Server) StartServer() {
	config, err := util.LoadConfig("./config/user")
	if err != nil {
		log.Fatalf("error loading config file: %v", err)
	}
	if err := http.ListenAndServe(":"+config.LisAddr, s.r); err != nil {
		log.Fatalf("error starting the server at port :%v", config.LisAddr)
	}
}
