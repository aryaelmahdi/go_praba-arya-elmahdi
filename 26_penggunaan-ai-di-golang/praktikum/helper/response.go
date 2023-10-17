package helper

import (
	"encoding/json"
	"net/http"
)

func SetResponse(writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	return err
}
