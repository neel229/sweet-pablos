package user

import (
	"encoding/json"
	"log"

	user "github.com/neel229/sweet-pablos/pkg/user/service"
)

type Serializer struct {
}

func (s *Serializer) Decode(input []byte) (*user.User, error) {
	user := &user.User{}
	if err := json.Unmarshal(input, &user); err != nil {
		log.Fatal("error unmarshaling data")
		return nil, err
	}
	return user, nil
}

func (s *Serializer) Encode(user *user.User) ([]byte, error) {
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal("error marshaling data")
		return nil, err
	}
	return data, nil
}
