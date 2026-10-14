package userhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type ChangeUserPaswwordRequest struct {
	Current_Password string `json:"current_password"`
	New_Password     string `json:"new_password"`
}

func (h *UserHandler) ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	var changeUserPassword ChangeUserPaswwordRequest
	if err := json.NewDecoder(r.Body).Decode(&changeUserPassword); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	if err := h.userService.ChangeUserPassword(ctx,userID, changeUserPassword.Current_Password, changeUserPassword.New_Password); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	fmt.Println("user password successfully change.")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "password changed successfully"})

}
