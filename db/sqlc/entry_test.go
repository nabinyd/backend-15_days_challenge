package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertIntoEntry(t *testing.T) {

	account1 := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    100,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	fmt.Println(account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
}
