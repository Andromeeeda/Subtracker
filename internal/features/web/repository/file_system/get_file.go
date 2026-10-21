package web_repository

import (
	"fmt"
	"os"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

func(r *WebRepository) GetFile(filePath string) ([]byte,error) {
	file,err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil,fmt.Errorf("get file: %s %w",filePath,common_errors.ErrNotFound)
		}

		return nil,fmt.Errorf("get file: %s %w",filePath,err)
	}

	return file,nil 
}