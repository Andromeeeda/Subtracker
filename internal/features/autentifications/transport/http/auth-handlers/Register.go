package authhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHAndler) Register(w http.ResponseWriter, r *http.Request) {

	var registerRequest RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusBadRequest)

		return
	}

	if err := h.authService.Register(registerRequest.Username, registerRequest.Password); err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusConflict)

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string {
		"successfully text": "succesfylly added",
	})

}
