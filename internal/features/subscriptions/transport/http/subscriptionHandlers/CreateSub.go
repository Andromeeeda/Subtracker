package subscriptionhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type CreateSubRequest struct {
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

type CreateSubResponce struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

func (h *SubHandler) CreateSub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	var createSubRequest CreateSubRequest

	if err := json.NewDecoder(r.Body).Decode(&createSubRequest); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)

		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	Subscription, err := h.subService.CreateSubscription(ctx,userID, createSubRequest.Title, createSubRequest.Cost, createSubRequest.Date)
	if err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		if errors.Is(err, common_errors.ErrConflict) {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusConflict)
		} else {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusInternalServerError)
		}

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
