package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthCheckHandler struct {

}

func (handler *HealthCheckHandler) Handle(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	successResponse, _ := json.Marshal(map[string]bool{"success": true})

	if _, err := res.Write(successResponse); err != nil {
		log.Panic(err)
	}
}

func CreateHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}