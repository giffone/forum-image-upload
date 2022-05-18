package app

import (
	"github.com/giffone/forum-image-upload/internal/adapters/repository"
	"github.com/giffone/forum-image-upload/internal/service"
	"github.com/giffone/forum-image-upload/internal/service/category"
)

func (a *App) category(repo repository.Repo) service.Category {
	return category.NewService(repo)
}
