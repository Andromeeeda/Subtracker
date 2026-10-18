package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Andromeeeda/Subtracker/internal/core/domain"
	"github.com/Andromeeeda/Subtracker/internal/core/transport/http/middleware"
)

type AllSubResponce struct {
	Subscriptions []*domain.Subscription `json:"subscriptions"`
}

func (h *SubHandler) GetAllSub(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	AllSub,err := h.subService.GetAllSub(ctx,userID)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}

	resp := &AllSubResponce {
		Subscriptions: AllSub,
	}

	fmt.Println("All subscriptions are shown.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
