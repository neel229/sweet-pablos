package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/neel229/sweet-pablos/util"
	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	tokenInterface, err := NewJWT(util.RandomString(32))
	require.NoError(t, err)

	userID := util.RandomInt(1, 99)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := tokenInterface.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := tokenInterface.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second*2)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second*2)

	require.NotZero(t, payload.ID)
}

func TestExpiredJWT(t *testing.T) {
	tokenInterface, err := NewJWT(util.RandomString(32))
	require.NoError(t, err)

	userID := util.RandomInt(1, 99)
	duration := time.Minute

	token, err := tokenInterface.CreateToken(userID, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := tokenInterface.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidTokenAlgo(t *testing.T) {
	payload, err := NewPayload(util.RandomInt(1, 99), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	tokenInterface, err := NewJWT(util.RandomString(32))
	require.NoError(t, err)

	payload, err = tokenInterface.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrInvalidToken.Error())
	require.Nil(t, payload)
}
