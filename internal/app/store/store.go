package store

// Store ...
type Store interface {
	Note() NoteRepository
	User() UserRepository
}
