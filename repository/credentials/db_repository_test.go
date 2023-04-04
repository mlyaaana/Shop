package credentials

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

	userId := uuid.NewString()
	username := "trulya_test"
	password := "secret"
	err = repo.Create(&domain.Credentials{UserId: userId, Username: username, Password: password})
	require.NoError(t, err)

}
