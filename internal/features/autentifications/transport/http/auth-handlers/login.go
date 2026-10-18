package authhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponce struct {
	Session_id string `json:"session_id"`
}

func (h *AuthHAndler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var loginRequest LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	session_id,err := h.authService.Login(ctx,loginRequest.Username,loginRequest.Password)
	if err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusUnauthorized)

		return
	}

	loginResponce := &LoginResponce{
		Session_id: session_id,
	}

	fmt.Println("The user has successfully logged into the account.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponce)
	
}
