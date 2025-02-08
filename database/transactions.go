package database

import (
	"context"
	"fmt"
)

// RunTx executes a function within a database transaction
func RunTx(ctx context.Context, fn func(q *Queries) error) error {
	pool := GetDBPool() // Get global database pool
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	q := GetQueries().WithTx(tx) // Create a new Queries instance for the transaction

	if err := fn(q); err != nil {
		tx.Rollback(ctx) // Rollback if any error occurs
		return err
	}

	return tx.Commit(ctx) // Commit transaction
}
