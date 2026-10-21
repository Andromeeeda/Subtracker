package web_http_webHandler

type WebService interface {
	GetMainPage() ([]byte,error)
}

type WebHttpHandler struct {
	webService WebService
}

func NewWebHttpHandler(webService WebService) *WebHttpHandler {
	return &WebHttpHandler{
		webService: webService,
	}
}



