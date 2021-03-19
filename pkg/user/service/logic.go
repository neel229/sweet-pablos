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

func (u *userService) Find(code uuid.UUID) (*User, error) {
	return u.userRepo.Find(code)
}

func (u *userService) Update(code uuid.UUID, data *User) (*User, error) {
	_, err := u.Find(code)
	if err != nil {
		log.Fatalf("error finding user with id: %v", code)
		return nil, err
	}
	return u.userRepo.Update(data)
}

func (u *userService) Delete(code uuid.UUID) error {
	_, err := u.Find(code)
	if err != nil {
		log.Fatal("cannot delete user which doesn't exist")
		return err
	}
	return u.userRepo.Delete(code)
}

func (u *userService) Login(email, password string) error {
	// TODO: hash the password and send it to repo login function
	return u.userRepo.Login(email, password)
}
