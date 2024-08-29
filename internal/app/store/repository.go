package store

import "github.com/andreykvetinsky/http-rest-api-notes/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	DeleteByEmail(string) error
}

type NoteRepository interface {
	Create(*model.Note) error
	Find(int) (*model.Note, error)
	FindAllNotesByUserID(int) ([]*model.Note, error)
	DeleteNote(int) error
	DeleteNotes(int) error
}
