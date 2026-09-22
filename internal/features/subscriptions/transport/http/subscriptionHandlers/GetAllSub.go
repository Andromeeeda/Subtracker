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

	userID, ok := middleware.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unathorized", http.StatusUnauthorized)
		return
	}

	AllSub := h.subService.GetAllSub(userID)

	fmt.Println("All subscriptions are shown.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AllSub)

}
