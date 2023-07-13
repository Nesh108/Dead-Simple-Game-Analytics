package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"gorm.io/gorm"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/services"
)

type controller struct {
	DB *gorm.DB
}

func New(db *gorm.DB) controller {
	return controller{db}
}

// use http statuses for different errors: https://go.dev/src/net/http/status.go
func (c controller) ValidationErrorResponse(w http.ResponseWriter, message string) {
	c.Response(w, message, http.StatusBadRequest)
}

func (c controller) UnhandledErrorResponse(w http.ResponseWriter, message string, err error) {
	fmt.Println(err)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "N/A"
	}

	errorMessage := fmt.Sprintf("[%s] DeadSimpleGameAnalytics: %s (%s).", hostname, message, err.Error())
	services.SendTelegramMessages(errorMessage)
	c.Response(w, errorMessage, http.StatusInternalServerError)
}

func (c controller) SuccessResponse(w http.ResponseWriter) {
	c.Response(w, "success", http.StatusOK)
}

func (c controller) Response(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var response models.Response
	response.Message = message
	encodeErr := json.NewEncoder(w).Encode(response)

	if encodeErr != nil && !errors.Is(encodeErr, syscall.EPIPE) {
		c.UnhandledErrorResponse(w, "Failed to encode Controller Response", encodeErr)
	}
}
