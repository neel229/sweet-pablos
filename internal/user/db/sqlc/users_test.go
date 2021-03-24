package db

import (
	"context"
	"testing"

	"github.com/neel229/sweet-pablos/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Uaccount {
	arg := CreateUserParams{
		FirstName: util.RandomString(5),
		LastName:  util.RandomString(5),
		Email:     util.RandomString(8) + "@gmail.com",
		Password:  util.RandomString(12),
	}
	account, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.FirstName, account.FirstName)
	require.Equal(t, arg.LastName, account.LastName)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Password, account.Password)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.Ts)

	return account
}

func TestCreateUAccount(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomUser(t)

	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)

	require.NotZero(t, account1.ID)
	require.NotZero(t, account1.Ts)

}

func TestGetAccountByEmail(t *testing.T) {
	account1 := createRandomUser(t)

	account2, err := testQueries.GetUserByEmail(context.Background(), account1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.FirstName, account2.FirstName)
	require.Equal(t, account1.LastName, account2.LastName)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)

	require.NotZero(t, account1.ID)
	require.NotZero(t, account1.Ts)

}

func TestListAccounts(t *testing.T) {
	var i int
	for i < 10 {
		createRandomUser(t)
		i++
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateFirstName(t *testing.T) {
	account1 := createRandomUser(t)
	arg := UpdateFirstNameParams{
		ID:        account1.ID,
		FirstName: util.RandomString(3),
	}
	err := testQueries.UpdateFirstName(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateLastName(t *testing.T) {
	account1 := createRandomUser(t)
	arg := UpdateLastNameParams{
		ID:       account1.ID,
		LastName: util.RandomString(3),
	}
	err := testQueries.UpdateLastName(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateEmail(t *testing.T) {
	account1 := createRandomUser(t)
	arg := UpdateEmailParams{
		ID:    account1.ID,
		Email: util.RandomString(8) + "@gmail.com",
	}
	err := testQueries.UpdateEmail(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	account1 := createRandomUser(t)
	arg := UpdatePasswordParams{
		ID:       account1.ID,
		Password: util.RandomString(12),
	}
	err := testQueries.UpdatePassword(context.Background(), arg)
	require.NoError(t, err)
}
