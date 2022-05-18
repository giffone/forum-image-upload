package app

import (
	"github.com/giffone/forum-image-upload/internal/adapters/repository"
	"github.com/giffone/forum-image-upload/internal/service"
	"github.com/giffone/forum-image-upload/internal/service/filemaker"
)

func (a *App) file(repo repository.Repo) service.FileMaker {
	srv := filemaker.NewService(repo)
	return srv
}
