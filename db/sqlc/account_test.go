package db

import (
	"context"
	"testing"

	"github.com/eltanvir/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRanndomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQuries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.Balance, account.Balance)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRanndomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account1 := CreateRanndomAccount(t)
	account2, err := testQuries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.ID, account2.ID)
	require.NotZero(t, account2.CreatedAt)

}
