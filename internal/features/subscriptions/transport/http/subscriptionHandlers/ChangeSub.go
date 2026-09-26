package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
	"github.com/gorilla/mux"
)

type ChangeSubRequest struct {
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

type ChangeSubResponce struct {
	ID    int     `json:"uuid"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

func (h *SubHandler) ChangeSub(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	var changeSub ChangeSubRequest
	if err := json.NewDecoder(r.Body).Decode(&changeSub); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	ChangeSub, err := h.subService.ChangeSub(userID,id, &changeSub.Title, &changeSub.Cost, &changeSub.Date)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusNotFound)

		return
	}

	changeSubResponce := &ChangeSubResponce{
		ID:    ChangeSub.ID,
		Title: ChangeSub.Title,
		Cost:  ChangeSub.Cost,
		Date:  ChangeSub.Date,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(changeSubResponce)

}
