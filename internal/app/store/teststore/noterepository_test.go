package teststore_test

import (
	"testing"

	"github.com/andreykvetinsky/http-rest-api/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api/internal/app/store"
	"github.com/andreykvetinsky/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNoteRepository_Create(t *testing.T) {
	s := teststore.New()
	n := model.TestNote(t)
	assert.NoError(t, s.Note().Create(n))
	assert.NotNil(t, n)
}

func TestNoteRepository_Find(t *testing.T) {
	s := teststore.New()
	n1 := model.TestNote(t)
	err := s.Note().Create(n1)
	assert.NoError(t, err)
	n2, err := s.Note().Find(n1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, n2)
}

func TestNoteRepository_FindAllNotesByUserID(t *testing.T) {

	s := teststore.New()
	user_id := 10
	_, err := s.Note().FindAllNotesByUserID(user_id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	n := model.TestNote(t)
	n.User_ID = user_id

	s.Note().Create(n)
	n1, err := s.Note().FindAllNotesByUserID(user_id)
	assert.NoError(t, err)
	assert.NotNil(t, n1)

}

func TestUserRepository_DeleteNote(t *testing.T) {
	s := teststore.New()

	n := model.TestNote(t)
	s.Note().Create(n)
	err := s.Note().DeleteNote(n.ID)
	assert.NoError(t, err)
	n1, err := s.Note().Find(n.ID)
	assert.Error(t, err)
	assert.Nil(t, n1)
}
