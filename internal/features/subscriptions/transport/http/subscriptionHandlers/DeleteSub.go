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

func (h *SubHandler) DeleteSub(w http.ResponseWriter, r *http.Request) {
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

	if err := h.subService.DeleteSub(id); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusNotFound)

		return
	}
	
	fmt.Println("subcription removed!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[int]string{})

}
