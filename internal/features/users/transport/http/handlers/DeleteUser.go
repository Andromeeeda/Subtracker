package userhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type DeleteUserRequest struct {
	Password string `json:"password"`
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	var deleteUser DeleteUserRequest

	if err := json.NewDecoder(r.Body).Decode(&deleteUser); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	if err := h.userService.DeleteUser(userID,deleteUser.Password); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}	

	fmt.Println("user successfully delete")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {"message": "user successfully delete"})

}
