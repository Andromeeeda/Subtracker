package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type CreateSubRequest struct {
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

type CreateSubResponce struct {
	ID    int     `json:"uuid"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

func (h *SubHandler) CreateSub(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w,"unathorized",http.StatusUnauthorized)
		return
	}

	var createSubRequest CreateSubRequest

	if err := json.NewDecoder(r.Body).Decode(&createSubRequest); err != nil {
		errDto := errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)

		http.Error(w, errors.ErrorsDtoToString(&errDto), http.StatusBadRequest)

		return
	}

	Subscription, err := h.subService.CreateSubscription(userID,createSubRequest.Title, createSubRequest.Cost, createSubRequest.Date)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	subResponce := CreateSubResponce{
		ID:    Subscription.ID,
		Title: Subscription.Title,
		Cost:  Subscription.Cost,
		Date:  Subscription.Date,
	}

	fmt.Println("subscription succsesfully added!")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subResponce)

}
