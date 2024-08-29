package sqlstore_test

import (
	"testing"

	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestNoteRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("notes")

	s := sqlstore.New(db)
	n := model.TestNote(t)
	err := s.Note().Create(n)
	assert.NoError(t, err)
	assert.NotNil(t, n)
	err = s.Note().DeleteNote(1)
	assert.NoError(t, err)
}

func TestNoteRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("notes")

	s := sqlstore.New(db)
	n := model.TestNote(t)
	err := s.Note().Create(n)
	assert.NoError(t, err)
	n2, err := s.Note().Find(n.ID)
	assert.NoError(t, err)
	assert.NotNil(t, n2)
	err = s.Note().DeleteNote(n.ID)
	assert.NoError(t, err)
}

func TestNoteRepository_FindAllNotesByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("notes")

	s := sqlstore.New(db)
	n := model.TestNote(t)
	err := s.Note().Create(n)
	assert.NoError(t, err)

	n1, err := s.Note().FindAllNotesByUserID(n.User_ID)
	assert.NoError(t, err)
	assert.NotNil(t, n1)

	err = s.Note().DeleteNotes(n.User_ID)
	assert.NoError(t, err)
}

func TestNoteRepository_DeleteNote(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("notes")

	s := sqlstore.New(db)
	n := model.TestNote(t)
	err := s.Note().Create(n)
	assert.NoError(t, err)

	err = s.Note().DeleteNote(n.ID)
	assert.NoError(t, err)
	_, err = s.Note().Find(n.ID)
	assert.Error(t, err)

}

func TestNoteRepository_DeleteNotes(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("notes")

	s := sqlstore.New(db)
	n := model.TestNote(t)
	err := s.Note().Create(n)
	assert.NoError(t, err)
	err = s.Note().DeleteNotes(n.User_ID)
	assert.NoError(t, err)
	_, err = s.Note().FindAllNotesByUserID(n.User_ID)
	assert.Error(t, err)
}
