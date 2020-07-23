package main

import (
	"github.com/ohatakky/dashboard/note/handler/cmd"
	"github.com/ohatakky/dashboard/note/repository"
	"github.com/ohatakky/dashboard/note/usecase"
	"github.com/ohatakky/dashboard/pkg/note"
	"github.com/ohatakky/dashboard/project/configs"
)

func main() {
	configs.InitConfigs()

	{
		client := note.NewClient(configs.E.Note.User)
		repo := repository.NewNoteRepository(client)
		uc := usecase.NewNoteUsecase(repo)
		cmdHandler := cmd.NewCmdNoteHandler(uc)
		cmdHandler.GetPosts()
	}
}
