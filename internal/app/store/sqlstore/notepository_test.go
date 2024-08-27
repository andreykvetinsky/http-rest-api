package sqlstore_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/andreykvetinsky/http-rest-api/internal/app/model"
// 	"github.com/andreykvetinsky/http-rest-api/internal/app/store"
// 	"github.com/andreykvetinsky/http-rest-api/internal/app/store/sqlstore"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNoteRepository_Create(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("notes")

// 	s := sqlstore.New(db)
// 	n := model.TestNote(t)
// 	fmt.Println(n, s)
// 	err := s.Note().Create(n)
// 	fmt.Println(n)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, n)
// 	err = s.Note().DeleteNote(1)
// 	assert.NoError(t, err)
// }

// func TestNoteRepository_Find(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")

// 	s := sqlstore.New(db)
// 	u1 := model.TestUser(t)
// 	err := s.User().Create(u1)
// 	assert.NoError(t, err)
// 	u2, err := s.User().Find(u1.ID)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u2)
// 	err = s.User().DeleteByEmail(u1.Email)
// 	assert.NoError(t, err)

// }

// func TestNoteRepository_FindAllNotesByUserID(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")

// 	s := sqlstore.New(db)
// 	email := "test126@example.com"
// 	_, err := s.User().FindByEmail(email)
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

// 	u := model.TestUser(t)
// 	s.User().Create(u)
// 	u1, err := s.User().FindByEmail(u.Email)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u1)
// 	err = s.User().DeleteByEmail(u.Email)
// 	assert.NoError(t, err)
// }

// func TestNoteRepository_DeleteNote(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")

// 	s := sqlstore.New(db)

// 	u := model.TestUser(t)
// 	s.User().Create(u)
// 	err := s.User().DeleteByEmail(u.Email)
// 	assert.NoError(t, err)
// 	u1, err := s.User().FindByEmail(u.Email)
// 	assert.Error(t, err)
// 	assert.Nil(t, u1)
// }
