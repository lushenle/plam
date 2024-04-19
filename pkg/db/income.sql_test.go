package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lushenle/plam/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomIncome(t *testing.T) Income {
	project := createRandomProject(t)
	require.NotEmpty(t, project)

	arg := CreateIncomeParams{
		Payee:     util.RandomString(10),
		Amount:    util.RandomFloat32(0, 100),
		ProjectID: uuid.MustParse(project.ID),
	}

	income, err := testStore.CreateIncome(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, income)

	require.Equal(t, arg.Payee, income.Payee)
	require.Equal(t, arg.Amount, income.Amount)
	require.Equal(t, arg.ProjectID, income.ProjectID)
	require.NotZero(t, income.CreatedAt)

	return income
}

func TestCreateIncome(t *testing.T) {
	createRandomIncome(t)
}

func TestGetIncome(t *testing.T) {
	income1 := createRandomIncome(t)

	income2, err := testStore.GetIncome(context.Background(), income1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, income2)

	require.Equal(t, income1.ID, income2.ID)
	require.Equal(t, income1.Payee, income2.Payee)
	require.Equal(t, income1.Amount, income2.Amount)
	require.Equal(t, income1.ProjectID, income2.ProjectID)
	require.WithinDuration(t, income1.CreatedAt, income2.CreatedAt, 0)
}

func TestListIncomes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIncome(t)
	}

	arg := ListIncomesParams{
		Limit:  5,
		Offset: 0,
	}
	incomes, err := testStore.ListIncomes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, incomes, 5)

	for _, income := range incomes {
		require.NotEmpty(t, income)
	}
}

func TestSearchIncomes(t *testing.T) {
	project := createRandomProject(t)
	require.NotEmpty(t, project)

	arg := CreateIncomeParams{
		Payee:     fmt.Sprintf("%s%s", "search", util.RandomString(10)),
		Amount:    util.RandomFloat32(0, 100),
		ProjectID: uuid.MustParse(project.ID),
	}

	income, err := testStore.CreateIncome(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, income)

	require.Equal(t, arg.Payee, income.Payee)
	require.Equal(t, arg.Amount, income.Amount)
	require.Equal(t, arg.ProjectID, income.ProjectID)
	require.NotZero(t, income.CreatedAt)

	arg2 := SearchIncomesParams{
		Column1: pgtype.Text{
			String: arg.Payee,
			Valid:  true,
		},
		Offset: 0,
		Limit:  5,
	}
	income2, err := testStore.SearchIncomes(context.Background(), arg2)
	require.NoError(t, err)
	require.Len(t, income2, 1)

	require.Equal(t, income.ID, income2[0].ID)
	require.Equal(t, income.Payee, income2[0].Payee)
	require.Equal(t, income.Amount, income2[0].Amount)
	require.Equal(t, income.ProjectID, income2[0].ProjectID)
	require.WithinDuration(t, income.CreatedAt, income2[0].CreatedAt, 0)
}

func TestDeleteIncome(t *testing.T) {
	income1 := createRandomIncome(t)

	income2, err := testStore.DeleteIncome(context.Background(), income1.ID)
	require.NoError(t, err)

	require.NotEmpty(t, income2)
	require.Equal(t, income1.ID, income2.ID)
	require.Equal(t, income1.Payee, income2.Payee)
	require.Equal(t, income1.Amount, income2.Amount)
	require.Equal(t, income1.ProjectID, income2.ProjectID)
	require.WithinDuration(t, income1.CreatedAt, income2.CreatedAt, 0)

	income3, err := testStore.GetIncome(context.Background(), income1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRecordNotFound)
	require.Empty(t, income3)
}
