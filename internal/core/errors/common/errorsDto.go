package common_errors

import (
	"encoding/json"
	"time"
)

type ErrorsDTO struct {
	Error string    `json:"error"`
	Time  time.Time `json:"time"`
}

func ErrorsDtoToString(errDto *ErrorsDTO) string {
	errDtoJson, _ := json.Marshal(errDto)

	return string(errDtoJson)
}
