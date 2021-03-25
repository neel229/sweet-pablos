package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	db "github.com/neel229/sweet-pablos/internal/user/db/sqlc"
	"github.com/neel229/sweet-pablos/util"
)

type createUserParams struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required, len>=8"`
}

// CreateUser endpoint is used for creating a new user
func (s *Server) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := new(createUserParams)
		json.NewDecoder(r.Body).Decode(&data)

		hashedPass, err := util.HashPassword(data.Password)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "there was an internal server error. Please try again.", http.StatusInternalServerError)
			return
		}

		arg := db.CreateUserParams{
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Password:  hashedPass,
		}

		v := validator.New()
		err = v.Struct(arg)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "invalid data provided", http.StatusBadRequest)
			return
		}

		user, err := s.store.CreateUser(context.Background(), arg)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "error storing user data", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(user)
	}
}

// GetUser endpoint retrieves the user with id provided in the URL
func (s *Server) GetUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		user, err := s.store.GetUser(context.Background(), id)
		if err != nil {
			log.Fatal(err)
			if err == sql.ErrNoRows {
				http.Error(rw, fmt.Sprintf("user with id: %v doesn't exists", id), http.StatusBadRequest)
				return
			}
			http.Error(rw, "Error creating user. Please try again later.", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(user)
	}
}

type updateEmailReqParams struct {
	Email string `json:"email" validate:"required,email"`
}

// UpdateEmail endpoint updates the email of the user
// whose id is provided in the url
func (s *Server) UpdateEmail() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := new(updateEmailReqParams)
		json.NewDecoder(r.Body).Decode(&data)

		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		arg := db.UpdateEmailParams{
			ID:    id,
			Email: data.Email,
		}

		v := validator.New()
		err := v.Struct(arg)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "invalid data provided", http.StatusBadRequest)
			return
		}
		if err := s.store.UpdateEmail(context.Background(), arg); err != nil {
			log.Fatal(err)
			http.Error(rw, "Error updating email. Please try again later", http.StatusInternalServerError)
			return
		}
	}
}

type updatePasswordReqParams struct {
	Password string `json:"password" validate:"required,len>=8"`
}

// UpdatePassword endpoint updates the password of the user
// whose id is provided in the url
func (s *Server) UpdatePassword() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := new(updatePasswordReqParams)
		json.NewDecoder(r.Body).Decode(&data)

		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		arg := db.UpdatePasswordParams{
			ID:       id,
			Password: data.Password,
		}

		v := validator.New()
		err := v.Struct(arg)
		if err != nil {
			log.Fatal(err)
			http.Error(rw, "password's min. length has to be 8 ", http.StatusBadRequest)
			return
		}
		if err := s.store.UpdatePassword(context.Background(), arg); err != nil {
			log.Fatal(err)
			http.Error(rw, "Error updating password. Please try again later", http.StatusInternalServerError)
			return
		}
	}
}

type updateFNameReqParams struct {
	FirstName string `json:"first_name"`
}

// UpdateFirstName endpoint updates the first name of the user
// whose id is provided in the url
func (s *Server) UpdateFirstName() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := new(updateFNameReqParams)
		json.NewDecoder(r.Body).Decode(&data)

		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		arg := db.UpdateFirstNameParams{
			ID:        id,
			FirstName: data.FirstName,
		}

		if err := s.store.UpdateFirstName(context.Background(), arg); err != nil {
			log.Fatal(err)
			http.Error(rw, "Error updating first name. Please try again later", http.StatusInternalServerError)
			return
		}
	}
}

type updateLNameReqParams struct {
	LastName string `json:"last_name"`
}

// UpdateLastName endpoint updates the last name of the user
// whose id is provided in the url
func (s *Server) UpdateLastName() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data := new(updateLNameReqParams)
		json.NewDecoder(r.Body).Decode(&data)

		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		arg := db.UpdateLastNameParams{
			ID:       id,
			LastName: data.LastName,
		}

		if err := s.store.UpdateLastName(context.Background(), arg); err != nil {
			log.Fatal(err)
			http.Error(rw, "Error updating last name. Please try again later", http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) DeleteUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)

		if err := s.store.DeleteUser(context.Background(), id); err != nil {
			log.Fatal(err)
			http.Error(rw, "Error deleting your account, please try again later.", http.StatusInternalServerError)
			return
		}
	}
}
