package main

import (
	"github.com/ohatakky/dashboard/server/note/handler/cmd"
	"github.com/ohatakky/dashboard/server/note/repository"
	"github.com/ohatakky/dashboard/server/note/usecase"
	"github.com/ohatakky/dashboard/server/pkg/note"
	"github.com/ohatakky/dashboard/server/project/configs"
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
