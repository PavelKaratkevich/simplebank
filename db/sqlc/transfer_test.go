package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccount, toAccount Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQuery.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotNil(t, transfer.ID)
	require.NotNil(t, transfer.CreatedAt)

	require.Equal(t, transfer.FromAccountID, fromAccount.ID)
	require.Equal(t, transfer.ToAccountID, toAccount.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {

	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {

	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)
	transfer1 := createRandomTransfer(t, fromAccount, toAccount)

	transfer2, err := testQuery.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQuery.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, j := range transfers {
		require.NotEmpty(t, j)
	}
}
