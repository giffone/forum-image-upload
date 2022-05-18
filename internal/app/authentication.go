package app

import (
	"github.com/giffone/forum-image-upload/internal/adapters/api"
	authentication2 "github.com/giffone/forum-image-upload/internal/adapters/api/authentication"
	"github.com/giffone/forum-image-upload/internal/adapters/authentication"
	"github.com/giffone/forum-image-upload/internal/service"
)

func (a *App) authentication(auth *authentication.Auth,
	srvUser service.User, sMid service.Middleware, aMid api.Middleware) {
	authentication2.NewHandler(auth, srvUser, sMid).Register(a.ctx, a.router, aMid)
}
