package db

import (
	"context"
	"testing"

	"github.com/mkarimi-s/bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer{
	randomAccount1 := CreateRandomAccount(t)
	randomAccount2 := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: randomAccount1.ID,
		ToAccountID: randomAccount2.ID,
		Amount: util.RandomMoney(),
	}
	
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer := CreateRandomTransfer(t)
	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, createdTransfer, transfer)
}