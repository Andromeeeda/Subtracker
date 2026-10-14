package userhandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
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

	ctx := r.Context()

	username, err := h.userService.GetUser(ctx,userID)
	if err != nil {
		errDto := &common_errors.ErrorsDTO{
			Error: err.Error(),
			Time:  time.Now(),
		}

		if errors.Is(err, common_errors.ErrNotFound) {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusNotFound)
		} else {
			fmt.Println("Error!", errDto)
			http.Error(w, common_errors.ErrorsDtoToString(errDto), http.StatusInternalServerError)
		}

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
