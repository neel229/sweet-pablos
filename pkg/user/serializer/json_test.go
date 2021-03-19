package user

import (
	"testing"

	user "github.com/neel229/sweet-pablos/pkg/user/service"
	"github.com/stretchr/testify/require"
)

func TestSerialzier(t *testing.T) {
	t.Parallel()

	var s Serializer
	user := user.User{
		Firstname: "abdcd",
		Lastname:  "ldsjfl",
		Email:     "abcd@gmail.com",
		Password:  "slkdfjas@5152",
		Address: user.Address{
			Street: "lorem",
			City:   "ipsum",
			State:  "dolor",
		},
	}
	data, err := s.Encode(&user)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	rUser, err := s.Decode(data)
	require.NoError(t, err)
	require.NotEmpty(t, rUser)

	require.Equal(t, rUser.Firstname, user.Firstname)
	require.Equal(t, rUser.Lastname, user.Lastname)
	require.Equal(t, rUser.Email, user.Email)
	require.Equal(t, rUser.Password, user.Password)
	require.Equal(t, rUser.Address, user.Address)

}
