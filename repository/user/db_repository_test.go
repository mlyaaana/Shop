package user

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"shop/database"
	"shop/domain"
	"testing"
)

func TestDBRepository(t *testing.T) {
	db, err := database.New()
	require.NoError(t, err)
	repo := NewDBRepository(db)

	users, err := repo.List()
	require.NoError(t, err)
	count := len(users)

	id := uuid.NewString()
	name := "test"
	err = repo.Create(&domain.User{Id: id, Name: name})
	require.NoError(t, err)

	user, err := repo.Get(id)
	require.NoError(t, err)
	require.Equal(t, name, user.Name)

	users, err = repo.List()
	require.NoError(t, err)
	require.Len(t, users, count+1)
	user = users[count]
	require.Equal(t, name, user.Name)

	repo.Delete(id)
	users, err = repo.List()
	require.NoError(t, err)
	require.Len(t, users, count)
	_, err = repo.Get(id)
	require.Error(t, err)
}
