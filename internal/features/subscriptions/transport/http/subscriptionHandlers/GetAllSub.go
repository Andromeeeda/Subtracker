package subscriptionhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Andromeeeda/Subtracker/internal/core/domain"
)


type AllSubResponce struct {
	Subscriptions []*domain.Subscription`json:"subscriptions"`
}

func (h *SubHandler) GetAllSub(w http.ResponseWriter, r *http.Request) {
	
	AllSub := h.subService.GetAllSub()	

	fmt.Println("All subscriptions are shown.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AllSub)

}
