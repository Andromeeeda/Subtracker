package authhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHAndler) Register(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var registerRequest RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	if err := h.authService.Register(ctx,registerRequest.Username, registerRequest.Password); err != nil {
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

	fmt.Println("The user has successfully registered.")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"successfully text": "succesfylly registered",
	})

}
