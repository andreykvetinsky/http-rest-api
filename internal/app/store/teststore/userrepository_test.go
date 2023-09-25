package teststore_test

import (
	"testing"

	"github.com/andreykvetinsky/http-rest-api/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api/internal/app/store"
	"github.com/andreykvetinsky/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {

	s := teststore.New()
	u1 := model.TestUser(t)
	err := s.User().Create(u1)
	assert.NoError(t, err)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "test66@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_DeleteByEmail(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	s.User().Create(u)
	err := s.User().DeleteByEmail(u.Email)
	assert.NoError(t, err)
	u1, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)
	assert.Nil(t, u1)
}
