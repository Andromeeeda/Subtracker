package main

import (
	"fmt"
	"net/http"

	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
	authservice "github.com/Andromeeeda/Subtracker/internal/features/autentifications/service/auth-service"
	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	authhandlers "github.com/Andromeeeda/Subtracker/internal/features/autentifications/transport/http/auth-handlers"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository/sub-rep"
	"github.com/Andromeeeda/Subtracker/internal/features/subscriptions/service"
	subscriptionhandlers "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/transport/http/subscriptionHandlers"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	userservice "github.com/Andromeeeda/Subtracker/internal/features/users/service/user_service"
	user_handlers "github.com/Andromeeeda/Subtracker/internal/features/users/transport/http/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Hello subtracker!")

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("starting subtracker server")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router := mux.NewRouter()
	router.Use(middleware.Logging(sugar))

	SubRep := subrep.NewSubRep()
	SubService := service.NewSubService(SubRep)
	SubHandler := subscriptionhandlers.NewSubHandler(SubService)

	SessionStore := sessionstorage.NewSessionStorage()

	UserRep := usersrepository.NewUsersRep()
	UserService := userservice.NewUserService(UserRep, SessionStore, SubRep)
	UserHandler := user_handlers.NewUserHandler(UserService)

	AuthService := authservice.NewAuthService(UserRep, SessionStore)
	AuthHandler := authhandlers.NewAuthHandler(AuthService)

	SubRouter := router.PathPrefix("/subscriptions").Subrouter()
	SubRouter.Use(mux.MiddlewareFunc(middleware.Auth(SessionStore)))

	SubRouter.Path("").Methods("POST").HandlerFunc(SubHandler.CreateSub)
	SubRouter.Path("").Methods("GET").HandlerFunc(SubHandler.GetAllSub)
	SubRouter.Path("/{id}").Methods("GET").HandlerFunc(SubHandler.GetSub)
	SubRouter.Path("/{id}").Methods("PATCH").HandlerFunc(SubHandler.ChangeSub)
	SubRouter.Path("/{id}").Methods("DELETE").HandlerFunc(SubHandler.DeleteSub)

	router.Path("/auth/register").Methods("POST").HandlerFunc(AuthHandler.Register)
	router.Path("/auth/login").Methods("POST").HandlerFunc(AuthHandler.Login)

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(mux.MiddlewareFunc(middleware.Auth(SessionStore)))

	userRouter.Path("/me").Methods("GET").HandlerFunc(UserHandler.GetUser)
	userRouter.Path("/me/password").Methods("POST").HandlerFunc(UserHandler.ChangeUserPassword)
	userRouter.Path("/me").Methods("DELETE").HandlerFunc(UserHandler.DeleteUser)

	if err := http.ListenAndServe(":8081", corsHandler(router)); err != nil {
		fmt.Println("fail to listen server")
	}

}
