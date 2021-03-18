package user

// UserSerializer has methods Encode and Decode attached to it
// for serializing data
type UserSerializer interface {
	Decode(input []byte) (*User, error)
	Encode(user *User) ([]byte, error)
}
