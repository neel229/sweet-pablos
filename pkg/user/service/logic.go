package user

import (
	"fmt"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Create(user *User) error {
	if err := validator.New().Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
	}
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("error generating uuid: %v", err)
		return err
	}
	user.UUID = uuid
	return u.userRepo.Create(user)
}
