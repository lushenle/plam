package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/lushenle/plam/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomPayOut(t *testing.T) PayOut {
	arg := CreatePayOutParams{
		Owner:   util.RandomString(10),
		Amount:  util.RandomFloat32(0, 100),
		Subject: util.RandomString(30),
	}
	payOut, err := testStore.CreatePayOut(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payOut)

	require.Equal(t, arg.Owner, payOut.Owner)
	require.Equal(t, arg.Amount, payOut.Amount)
	require.Equal(t, arg.Subject, payOut.Subject)
	require.NotZero(t, payOut.CreatedAt)

	return payOut
}

func TestCreatePayOut(t *testing.T) {
	createRandomPayOut(t)
}

func TestGetPayOut(t *testing.T) {
	payOut1 := createRandomPayOut(t)

	payOut2, err := testStore.GetPayOut(context.Background(), payOut1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payOut2)

	require.Equal(t, payOut1.ID, payOut2.ID)
	require.Equal(t, payOut1.Owner, payOut2.Owner)
	require.Equal(t, payOut1.Amount, payOut2.Amount)
	require.Equal(t, payOut1.Subject, payOut2.Subject)
	require.WithinDuration(t, payOut1.CreatedAt, payOut2.CreatedAt, 0)
}

func TestListPayOuts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPayOut(t)
	}

	arg := ListPayOutsParams{
		Limit:  5,
		Offset: 0,
	}
	payOuts, err := testStore.ListPayOuts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, payOuts, 5)

	for _, payOut := range payOuts {
		require.NotEmpty(t, payOut)
	}
}

func TestSearchPayOut(t *testing.T) {
	arg := CreatePayOutParams{
		Owner:   fmt.Sprintf("search-%s", util.RandomString(10)),
		Amount:  util.RandomFloat32(0, 100),
		Subject: util.RandomString(30),
	}
	payOut1, err := testStore.CreatePayOut(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payOut1)
	require.Equal(t, arg.Owner, payOut1.Owner)
	require.Equal(t, arg.Amount, payOut1.Amount)
	require.Equal(t, arg.Subject, payOut1.Subject)
	require.NotZero(t, payOut1.CreatedAt)

	arg2 := SearchPayOutsParams{
		Owner:  arg.Owner,
		Offset: 0,
		Limit:  5,
	}
	payOuts, err := testStore.SearchPayOuts(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, payOuts)

	for _, payOut := range payOuts {
		require.NotEmpty(t, payOut)
		require.Equal(t, arg.Owner, payOut.Owner)
		require.Equal(t, arg.Amount, payOut.Amount)
		require.Equal(t, arg.Subject, payOut.Subject)
		require.WithinDuration(t, payOut1.CreatedAt, payOut.CreatedAt, 0)
	}
}

func TestDeletePayOut(t *testing.T) {
	payOut1 := createRandomPayOut(t)

	payOut2, err := testStore.DeletePayOut(context.Background(), payOut1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payOut2)
	require.Equal(t, payOut1.ID, payOut2.ID)

	require.Equal(t, payOut1.Owner, payOut2.Owner)
	require.Equal(t, payOut1.Amount, payOut2.Amount)
	require.Equal(t, payOut1.Subject, payOut2.Subject)
	require.WithinDuration(t, payOut1.CreatedAt, payOut2.CreatedAt, 0)

	payOut3, err := testStore.GetPayOut(context.Background(), payOut1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRecordNotFound)
	require.Empty(t, payOut3)
}
