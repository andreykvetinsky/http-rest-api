package sqlstore_test

import (
	"testing"

	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/store"
	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	err = s.User().DeleteByEmail(u.Email)
	assert.NoError(t, err)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	err := s.User().Create(u1)
	assert.NoError(t, err)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	err = s.User().DeleteByEmail(u1.Email)
	assert.NoError(t, err)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "test126@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	s.User().Create(u)
	u1, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u1)
	err = s.User().DeleteByEmail(u.Email)
	assert.NoError(t, err)
}

func TestUserRepository_DeleteByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	s.User().Create(u)
	err := s.User().DeleteByEmail(u.Email)
	assert.NoError(t, err)
	u1, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)
	assert.Nil(t, u1)
}
