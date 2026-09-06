package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Prime-code-max/TodoApp/internal/core/domain"
	core_http_server "github.com/Prime-code-max/TodoApp/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UserHTTPHandler {
	return &UserHTTPHandler{
		usersService: usersService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
