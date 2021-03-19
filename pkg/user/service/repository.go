package user

import "github.com/google/uuid"

// UserRepository has attached methods for performing CRUD operations
// on User struct
type UserRepository interface {
	Create(user *User) error
	Find(code uuid.UUID) (*User, error)
	Update(user *User) (*User, error)
	Delete(code uuid.UUID) error
	Login(email, password string) error
}
