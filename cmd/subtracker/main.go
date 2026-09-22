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
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello subtracker!")

	router := mux.NewRouter()

	SubRep := subrep.NewSubRep()
	SubService := service.NewSubService(SubRep)
	SubHandler := subscriptionhandlers.NewSubHandler(SubService)

	UserRep := usersrepository.NewUsersRep() 

	SessionStore := sessionstorage.NewSessionStorage()

	AuthService := authservice.NewAuthService(UserRep,SessionStore)
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

	if err := http.ListenAndServe(":8081", router); err != nil {
		fmt.Println("fail to listen server")
	}

}
