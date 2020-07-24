package repository

import "github.com/ohatakky/dashboard/server/pkg/note"

type NoteRepository interface {
	GetPosts() (*note.Posts, error)
}

type noteRepository struct {
	note *note.Client
}

func NewNoteRepository(nc *note.Client) NoteRepository {
	return &noteRepository{
		note: nc,
	}
}

func (repo *noteRepository) GetPosts() (*note.Posts, error) {
	return repo.note.GetPosts()
}
