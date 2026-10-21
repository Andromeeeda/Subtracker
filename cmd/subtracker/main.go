package main

import (
	"context"
	"fmt"
	"net/http"

	core_postgres_pool "github.com/Andromeeeda/Subtracker/internal/core/repository/postgres/pool"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
	authservice "github.com/Andromeeeda/Subtracker/internal/features/autentifications/service/auth-service"
	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	authhandlers "github.com/Andromeeeda/Subtracker/internal/features/autentifications/transport/http/auth-handlers"
	subscription_repository_postgres "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository/postgres"
	"github.com/Andromeeeda/Subtracker/internal/features/subscriptions/service"
	subscriptionhandlers "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/transport/http/subscriptionHandlers"
	users_repository_postgres "github.com/Andromeeeda/Subtracker/internal/features/users/repository/postgres"
	userservice "github.com/Andromeeeda/Subtracker/internal/features/users/service/user_service"
	user_handlers "github.com/Andromeeeda/Subtracker/internal/features/users/transport/http/handlers"
	web_repository "github.com/Andromeeeda/Subtracker/internal/features/web/repository/file_system"
	web_service "github.com/Andromeeeda/Subtracker/internal/features/web/service"
	web_http_webHandler "github.com/Andromeeeda/Subtracker/internal/features/web/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Hello subtracker!")

	ctx := context.Background()

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

	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		sugar.Fatal("failet to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	SubRep := subscription_repository_postgres.NewSubPostgresRep(pool)
	SubService := service.NewSubService(SubRep)
	SubHandler := subscriptionhandlers.NewSubHandler(SubService)

	SessionStore := sessionstorage.NewSessionStorage()

	UserRep := users_repository_postgres.NewUsersPostgresRep(pool)
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

	WebRepository := web_repository.NewWebRepository()
	WebService := web_service.NewWebService(WebRepository)
	WebHandler := web_http_webHandler.NewWebHttpHandler(WebService)
	router.Path("/").Methods("GET").HandlerFunc(WebHandler.GetMainPage)

	if err := http.ListenAndServe(":8081", corsHandler(router)); err != nil {
		fmt.Println("fail to listen server")
	}

}
