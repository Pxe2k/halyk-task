package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func ToJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}

func ToErr(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ToJSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	ToJSON(w, http.StatusBadRequest, nil)
}
