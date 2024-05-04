package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/lushenle/plam/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomProject(t *testing.T) Project {
	arg := CreateProjectParams{
		Name:        util.RandomString(6),
		Description: util.RandomString(30),
		Amount:      util.RandomFloat32(0, 1000),
	}

	project, err := testStore.CreateProject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, project)

	require.Equal(t, arg.Name, project.Name)
	require.Equal(t, arg.Description, project.Description)
	require.Equal(t, arg.Amount, project.Amount)

	require.NotZero(t, project.CreatedAt)

	return project
}

func TestCreateProject(t *testing.T) {
	createRandomProject(t)
}

func TestGetProject(t *testing.T) {
	project1 := createRandomProject(t)

	project2, err := testStore.GetProject(context.Background(), project1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, project2)

	require.Equal(t, project1.ID, project2.ID)
	require.Equal(t, project1.Name, project2.Name)
	require.Equal(t, project1.Description, project2.Description)
	require.Equal(t, project1.Amount, project2.Amount)
	require.WithinDuration(t, project1.CreatedAt, project2.CreatedAt, 0)
}

func TestListProjects(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProject(t)
	}

	arg := ListProjectsParams{
		Limit:  5,
		Offset: 0,
	}
	projects, err := testStore.ListProjects(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, projects)
	require.Len(t, projects, 5)

	for _, project := range projects {
		require.NotEmpty(t, project)
	}
}

func TestSearchProjects(t *testing.T) {
	arg := CreateProjectParams{
		Name:        fmt.Sprintf("%s%s", "testpro", util.RandomString(8)),
		Description: util.RandomString(30),
		Amount:      util.RandomFloat32(300, 1000),
	}

	project, err := testStore.CreateProject(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, project)

	require.Equal(t, arg.Name, project.Name)
	require.Equal(t, arg.Description, project.Description)
	require.Equal(t, arg.Amount, project.Amount)

	require.NotZero(t, project.CreatedAt)

	searchArg := SearchProjectsParams{
		Name:   arg.Name,
		Offset: 0,
		Limit:  5,
	}
	result, err := testStore.SearchProjects(context.Background(), searchArg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	for _, p := range result {
		require.NotEmpty(t, p)
	}
}

func TestDeleteProject(t *testing.T) {
	project1 := createRandomProject(t)
	require.NotEmpty(t, project1)

	project2, err := testStore.DeleteProject(context.Background(), project1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, project2)

	require.Equal(t, project1.ID, project2.ID)
	require.Equal(t, project1.Name, project2.Name)
	require.Equal(t, project1.Description, project2.Description)
	require.Equal(t, project1.Amount, project2.Amount)
	require.WithinDuration(t, project1.CreatedAt, project2.CreatedAt, 0)

	project3, err := testStore.GetProject(context.Background(), project1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRecordNotFound)
	require.Empty(t, project3)
}
