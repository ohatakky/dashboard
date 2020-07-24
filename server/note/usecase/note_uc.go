package usecase

import (
	"github.com/ohatakky/dashboard/server/note/repository"
	"github.com/ohatakky/dashboard/server/pkg/note"
)

type NoteUsecase interface {
	GetPosts() (*note.Posts, error)
}

type noteUsecase struct {
	noteRepo repository.NoteRepository
}

func NewNoteUsecase(ar repository.NoteRepository) NoteUsecase {
	return &noteUsecase{
		noteRepo: ar,
	}
}

func (uc *noteUsecase) GetPosts() (*note.Posts, error) {
	return uc.noteRepo.GetPosts()
}
