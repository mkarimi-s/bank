package db

import (
	"context"
	"testing"

	"github.com/mkarimi-s/bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T){
	args := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, args.Owner)
	require.Equal(t, account.Balance, args.Balance)
	require.Equal(t, account.Currency, args.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}