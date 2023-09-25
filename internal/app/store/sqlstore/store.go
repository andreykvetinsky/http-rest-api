package sqlstore

import (
	"database/sql"

	"github.com/andreykvetinsky/http-rest-api/internal/app/store"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

// Store
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New creates a new Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
