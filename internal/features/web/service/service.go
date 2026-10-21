package web_service

type WebRepository interface {
	GetFile(filePath string) ([]byte,error) 
}

type WebService struct {
	webRepository WebRepository
}

func NewWebService(webRepository WebRepository) *WebService {
	return &WebService{
		webRepository: webRepository,
	}
} 
