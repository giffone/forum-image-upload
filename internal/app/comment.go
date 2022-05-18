package app

import (
	"github.com/giffone/forum-image-upload/internal/adapters/repository"
	"github.com/giffone/forum-image-upload/internal/service"
	"github.com/giffone/forum-image-upload/internal/service/comment"
)

func (a *App) comment(repo repository.Repo, srvRatio service.Ratio, sMid service.Middleware) service.Comment {
	srv := comment.NewService(repo, srvRatio, sMid)
	return srv
}
