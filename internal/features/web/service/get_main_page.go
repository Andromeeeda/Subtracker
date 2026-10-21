package web_service

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *WebService) GetMainPage() ([]byte, error) {
	root := os.Getenv("PROJECT_ROOT")
	if root == "" {
		root, _ = os.Getwd()
	}

	htmlFilePath := filepath.Join(root, "public", "index.html")

	html, err := s.webRepository.GetFile(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("read file from repository: %w", err)
	}

	return html, nil
}
