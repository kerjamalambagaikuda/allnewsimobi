package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	sqlc "github.com/kerjamalambagaikuda/allnewsimobi/db/sqlc"
	"github.com/stretchr/testify/require"
)

func createRandomStatus(t *testing.T) int64 {

	arg := sqlc.CreateStatusParams{
		Type:        "user",
		Code:        "active",
		Name:        "user active",
		Description: sql.NullString{String: "user is active", Valid: true},
	}

	status, err := testQueries.CreateStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, status)

	statusID, err := status.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(statusID)

	return statusID
}

func TestSelectStatusById(t *testing.T) {
	store := sqlc.NewStore(testDB)

	// statusID := createRandomStatus(t)

	// Cara tanpa menggunakan database transaction
	// status, err := testQueries.GetStatusById(context.Background(), statusID);

	// Cara dengan menggunakan database transaction
	status, err := store.GetStatusById(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, status)

	fmt.Println(status)
}
