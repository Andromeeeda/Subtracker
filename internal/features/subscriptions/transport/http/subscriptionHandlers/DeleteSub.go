package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
	"github.com/gorilla/mux"
)

func (h *SubHandler) DeleteSub(w http.ResponseWriter, r *http.Request) {
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

	if err := h.subService.DeleteSub(ctx,userID,id); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusNotFound)

		return
	}

	fmt.Println("subcription removed!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[int]string{})

}
