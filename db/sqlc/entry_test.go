package db

import (
	"context"
	"database/sql"
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

func TestDeleteEntry(t *testing.T) {
	createdEntry := CreateRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)

	deletedEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedEntry)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10 ; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit: 10,
		Offset: 0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)
	
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
} 
