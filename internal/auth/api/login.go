package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/neel229/sweet-pablos/util"
)

type loginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required, len>=8"`
}

func (s *Server) LoginUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req := new(loginReq)
		json.NewDecoder(r.Body).Decode(&req)
		pass, err := util.HashPassword(req.Password)
		if err != nil {
			log.Fatal(err)
			return
		}

		res, err := http.Get("http://localhost:9000/user?email=" + req.Email + "&password" + pass)
		if err != nil {
			log.Fatal(err)
			return
		}
		id, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		userID, _ := strconv.Atoi(string(id))
		token, err := s.t.CreateToken(int64(userID), s.c.AccessTokenDuration)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "error creating token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(token)
		log.Print("user logged in")
	}
}
