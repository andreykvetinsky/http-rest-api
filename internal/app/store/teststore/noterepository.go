package teststore

import (
	"errors"

	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api-notes/internal/app/store"
)

// UserRepository ...
type NoteRepository struct {
	store *Store
	notes map[int]*model.Note
}

// Create ...
func (r *NoteRepository) Create(n *model.Note) error {
	n.ID = len(r.notes) + 1
	r.notes[n.ID] = n

	return nil
}

// Find ...
func (r *NoteRepository) Find(id int) (*model.Note, error) {
	n, ok := r.notes[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return n, nil
}

// FindAllNotesByUserID...
func (r *NoteRepository) FindAllNotesByUserID(id int) ([]*model.Note, error) {
	notes := make([]*model.Note, 0)
	for _, n := range r.notes {
		if n.User_ID == id {
			notes = append(notes, n)
		}
	}
	if len(notes) == 0 {
		return nil, store.ErrRecordNotFound
	}
	return notes, nil
}

// DeleteNote...
func (r *NoteRepository) DeleteNote(id int) error {
	///....
	var err error = errors.New("not found id or incorrect id")
	// for _, n := range r.notes {
	_, ok := r.notes[id]
	if ok {
		delete(r.notes, id)
		return nil
	}
	// }
	return err
}

// DeleteNote...
func (r *NoteRepository) DeleteNotes(user_id int) error {
	///....
	var err error = errors.New("not found id or incorrect id")
	// for _, n := range r.notes {
	if r.notes != nil {
		r.notes = nil
		return nil
	}
	// }
	return err
}
