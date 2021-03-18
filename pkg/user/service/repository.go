package user

// UserRepository has attached methods for performing CRUD operations
// on User struct
type UserRepository interface {
	Create(user *User) error
	Find(code int) (*User, error)
	Update(user *User) (*User, error)
	Delete(code int) error
}
