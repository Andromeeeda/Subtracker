package userhandlers

type UserService interface {
	GetUser(userID string) (string, error)
	ChangeUserPassword(userID string,currentPassword string,newPassword string) error
	DeleteUser(userID string,password string) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
