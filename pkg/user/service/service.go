package user

import "github.com/google/uuid"

// UserService has attached methods for performing CRUD operations
// on User struct
type UserService interface {
	Create(user *User) error
	Find(code uuid.UUID) (*User, error)
	Update(code uuid.UUID, user *User) (*User, error)
	Delete(code uuid.UUID) error
	Login(user *User) (*User, error)
}
