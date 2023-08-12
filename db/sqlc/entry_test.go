package db

import (
	"context"
	"testing"
	"time"

	"github.com/mkarimi-s/bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry{
	account := CreateRandomAccount(t)
	require.NotEmpty(t, account)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomInt(0, 1000),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.WithinDuration(t, entry.CreatedAt, time.Now(), time.Second)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	createdEntry := CreateRandomEntry(t)

	entry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, createdEntry, entry)
}

func TestUpdateEntry(t *testing.T) {
	createdEntry := CreateRandomEntry(t)
	util.VarDump(createdEntry)

	arg := UpdateEntryParams{
		ID: createdEntry.ID,
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, arg.Amount, updatedEntry.Amount)
	require.Equal(t, createdEntry.AccountID, updatedEntry.AccountID)
}
