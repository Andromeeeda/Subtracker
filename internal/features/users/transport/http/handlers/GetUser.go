package userhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Andromeeeda/Subtracker/internal/core/errors"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type GetUserResponce struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	username, err := h.userService.GetUser(userID)
	if err != nil {
		errDto := &errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		fmt.Println("Error!", errDto)
		http.Error(w, errors.ErrorsDtoToString(errDto), http.StatusUnauthorized)

		return
	}

	responce := &GetUserResponce{
		UserID:   userID,
		Username: username,
	}

	fmt.Println("User data is successfully visible.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)

}
