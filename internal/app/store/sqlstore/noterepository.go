package sqlstore

import (
	"database/sql"

	"github.com/andreykvetinsky/http-rest-api/internal/app/model"
	"github.com/andreykvetinsky/http-rest-api/internal/app/store"
)

// UserRepository...
type NoteRepository struct {
	store *Store
}

// Create ...
func (r *NoteRepository) Create(n *model.Note) error {
	// if err := u.Validate(); err != nil {
	// 	return err
	// }

	if err := n.ValidateBeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO notes (user_id, note) VALUES ($1, $2) RETURNING id",
		n.User_ID,
		n.Note,
	).Scan(&n.ID)
}

func (r *NoteRepository) Find(id int) (*model.Note, error) {
	n := &model.Note{}

	if err := r.store.db.QueryRow(
		"SELECT id, user_id, note  FROM users WHERE id = $1",
		id,
	).Scan(
		&n.ID,
		&n.User_ID,
		&n.Note,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return n, nil
}

func (r *NoteRepository) FindAllNotesByUserID(id int) ([]*model.Note, error) {
	notes := make([]*model.Note, 0)

	rows, err := r.store.db.Query(
		"SELECT id, user_id, note  FROM notes WHERE user_id = $1",
		id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	for rows.Next() {
		n := &model.Note{}
		err := rows.Scan(&n.ID, &n.User_ID, &n.Note)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *NoteRepository) DeleteNote(id int) error {
	return r.store.db.QueryRow(
		"DELETE FROM notes WHERE id = $1",
		id,
	).Err()
}
