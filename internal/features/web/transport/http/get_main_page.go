package web_http_webHandler

import (
	"errors"
	"fmt"
	"net/http"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

func (h *WebHttpHandler) GetMainPage(w http.ResponseWriter, r *http.Request) {
	html,err := h.webService.GetMainPage()
	if err != nil {
		if errors.Is(err,common_errors.ErrNotFound) {
			http.Error(w,err.Error(),http.StatusNotFound)
			return 
		} else {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _,err := w.Write(html); err != nil {
		fmt.Println("write main page: %w",err)
	}
	
}
