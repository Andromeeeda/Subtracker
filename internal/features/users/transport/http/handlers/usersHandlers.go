package userhandlers

import "context"

type UserService interface {
	GetUser(ctx context.Context, userID string) (string, error)
	ChangeUserPassword(ctx context.Context,userID string, currentPassword string, newPassword string) error
	DeleteUser(ctx context.Context,userID string, password string) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
