package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Status(w http.ResponseWriter, statusCode int, msg any) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Status(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	Status(w, http.StatusBadRequest, nil)
}
