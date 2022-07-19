package db

import (
	"context"
	"database/sql"
	"fmt"
)

// We extend the struct to have it used not only by individual sql queries, but also by transactions. Therefore
// we embed *Queries into a new Store struct
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(q *Queries) error) error {

	// start a transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// create a new query with the help of the just created transaction tx
	q := New(tx)

	// apply callback function to a created query
	err = fn(q)
	if err != nil {
		// if there is an error, we roll back the transaction
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error %v", err, rbErr)
		}
		return err
	}
	// if no error, then we commit the transaction and return nil error
	return tx.Commit()
}

// TransferTxParams contains all necessary params for the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResults contains the results of the transfer transaction
type TransferTxResults struct {
	Transfer    Transfer `json:"tansfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

/* TransferTx performs the money transfer from one account to the other.
It creates a transfer record, two entries, and update accounts' balances within one DB transaction */
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var results TransferTxResults

	// execute transaction with callback function that create transfer
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// create trasfer and pass it to results
		results.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		results.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		results.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// in order to avoid deadlock, we process the queries one by one in function of their IDs
		if arg.FromAccountID < arg.ToAccountID {
			results.FromAccount, results.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if err != nil {
				return err
			}
			
		} else {
			results.FromAccount, results.ToAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return results, err
}

func addMoney(
	ctx context.Context, 
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	fromAccount, err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount:	amount1,
		ID: accountID1,
	})
	if err != nil {
		return
	}

	toAccount, err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount:	amount2,
		ID: accountID2,
	})
	if err != nil {
		return
	}
	return fromAccount, toAccount, err
}