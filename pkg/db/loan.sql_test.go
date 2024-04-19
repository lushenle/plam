package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lushenle/plam/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomLoan(t *testing.T) Loan {
	arg := CreateLoanParams{
		Borrower: util.RandomString(10),
		Amount:   util.RandomFloat32(0, 100),
		Subject:  util.RandomString(30),
	}
	loan, err := testStore.CreateLoan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, loan)

	require.Equal(t, arg.Borrower, loan.Borrower)
	require.Equal(t, arg.Amount, loan.Amount)
	require.Equal(t, arg.Subject, loan.Subject)
	require.NotZero(t, loan.CreatedAt)

	return loan
}

func TestCreateLoan(t *testing.T) {
	createRandomLoan(t)
}

func TestGetLoan(t *testing.T) {
	loan1 := createRandomLoan(t)

	loan2, err := testStore.GetLoan(context.Background(), loan1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, loan2)

	require.Equal(t, loan1.ID, loan2.ID)
	require.Equal(t, loan1.Borrower, loan2.Borrower)
	require.Equal(t, loan1.Amount, loan2.Amount)
	require.Equal(t, loan1.Subject, loan2.Subject)
	require.WithinDuration(t, loan1.CreatedAt, loan2.CreatedAt, 0)
}

func TestListLoans(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomLoan(t)
	}

	arg := ListLoansParams{
		Limit:  5,
		Offset: 0,
	}
	loans, err := testStore.ListLoans(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, loans, 5)
}

func TestSearchLoan(t *testing.T) {
	arg := CreateLoanParams{
		Borrower: fmt.Sprintf("search-%s", util.RandomString(10)),
		Amount:   util.RandomFloat32(0, 100),
		Subject:  util.RandomString(30),
	}
	loan, err := testStore.CreateLoan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, loan)

	require.Equal(t, arg.Borrower, loan.Borrower)
	require.Equal(t, arg.Amount, loan.Amount)
	require.Equal(t, arg.Subject, loan.Subject)
	require.NotZero(t, loan.CreatedAt)

	arg2 := SearchLoansParams{
		Column1: pgtype.Text{
			String: arg.Borrower,
			Valid:  true,
		},
		Offset: 0,
		Limit:  5,
	}
	loans, err := testStore.SearchLoans(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, loans)
	require.Equal(t, arg.Borrower, loans[0].Borrower)
	require.Equal(t, arg.Amount, loans[0].Amount)
	require.Equal(t, arg.Subject, loans[0].Subject)
	require.NotZero(t, loans[0].CreatedAt)
}

func TestDeleteLoan(t *testing.T) {
	loan1 := createRandomLoan(t)

	loan2, err := testStore.DeleteLoan(context.Background(), loan1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, loan2)
	require.Equal(t, loan1.ID, loan2.ID)
	require.Equal(t, loan1.Borrower, loan2.Borrower)
	require.Equal(t, loan1.Amount, loan2.Amount)
	require.Equal(t, loan1.Subject, loan2.Subject)
	require.WithinDuration(t, loan1.CreatedAt, loan2.CreatedAt, 0)

	loan3, err := testStore.GetLoan(context.Background(), loan1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRecordNotFound)
	require.Empty(t, loan3)
}
