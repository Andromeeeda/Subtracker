package main

import (
	"fmt"
	"net/http"

	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository/sub-rep"
	"github.com/Andromeeeda/Subtracker/internal/features/subscriptions/service"
	subscriptionhandlers "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/transport/http/subscriptionHandlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello subtracker!")

	router := mux.NewRouter()

	SubRep := subrep.NewSubRep()
	SubService := service.NewSubService(SubRep)
	SubHandler := subscriptionhandlers.NewSubHandler(SubService)

	router.Path("/subscriptions").Methods("POST").HandlerFunc(SubHandler.CreateSub)
	router.Path("/subscriptions").Methods("GET").HandlerFunc(SubHandler.GetAllSub)

	if err := http.ListenAndServe(":8081", router); err != nil {
		fmt.Println("fail to listen server")
	}

}
