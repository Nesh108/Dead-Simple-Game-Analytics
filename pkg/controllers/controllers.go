package controllers

import (
    "gorm.io/gorm"
    "encoding/json"
    "net/http"
    "fmt"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
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
    return
}

func (c controller) UnhandledErrorResponse(w http.ResponseWriter, err error) {
    fmt.Println(err)
    c.Response(w, err.Error(), http.StatusInternalServerError)
    return
}

func (c controller) SuccessResponse(w http.ResponseWriter) {
    c.Response(w, "success", http.StatusOK)
    return
}

func (c controller) Response(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    var response models.Response
    response.Message = message
    json.NewEncoder(w).Encode(response)
    return
}