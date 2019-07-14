package function

import (
	"encoding/json"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Handle a serverless request
func Handle(req []byte) string {
	resp := response{
		Success: true,
		Message: "hello world",
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	return string(respBytes)
}
