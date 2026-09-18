package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
	"github.com/gorilla/mux"
)

type GetSubResponce struct {
	ID    int     `json:"uuid"`
	Title string  `json:"title"`
	Cost  float64 `json:"cost"`
	Date  int     `json:"date"`
}

func (h *SubHandler) GetSub(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!",errDto)
		http.Error(w,errors.ErrorsDtoToString(errDto),http.StatusBadRequest)

		return 
	}

	sub,err :=h.subService.GetSub(id)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!",errDto)
		http.Error(w,errors.ErrorsDtoToString(errDto),http.StatusNotFound)

		return 
	}

	subResponce := GetSubResponce{
		ID: sub.ID,
		Title: sub.Title,
		Cost: sub.Cost,
		Date:sub.Date,
	}

	fmt.Println("subscription showed")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subResponce)

}
