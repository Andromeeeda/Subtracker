package subscriptionhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
	"github.com/gorilla/mux"
)

type GetSubResponce struct {
	ID    int     `json:"uuid"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

func (h *SubHandler) GetSub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	sub, err := h.subService.GetSub(ctx, userID, id)
	if err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		if errors.Is(err, common_errors.ErrNotFound) {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusNotFound)
		} else {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusInternalServerError)
		}

		return
	}

	subResponce := GetSubResponce{
		ID:    sub.ID,
		Title: sub.Title,
		Cost:  sub.Cost,
		Date:  sub.Date,
	}

	fmt.Println("subscription showed")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subResponce)

}
