package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}
	
	entry, err := testQuery.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}


func TestCreateEntry(t *testing.T) {

	account := CreateRandomAccount(t)
	CreateRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account)

	entry2, err := testQuery.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestListEntries(t *testing.T) {

	account := CreateRandomAccount(t)
	arg1 := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	for i := 0; i < 10; i++ {
		testQuery.CreateEntry(context.Background(), arg1)
	}

	arg2 := ListEntriesParams{
		AccountID: account.ID,
		Limit: 5, 
		Offset: 5,
	}

	entries, err := testQuery.ListEntries(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, k := range entries {
		require.NotEmpty(t, k)
	}
}