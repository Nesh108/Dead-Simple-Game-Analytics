package controllers

import (
    "gorm.io/gorm"
    "encoding/json"
    "net/http"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

type controller struct {
    DB *gorm.DB
}

func New(db *gorm.DB) controller {
    return controller{db}
}

func (c controller) ErrorResponse(w http.ResponseWriter, message string) {
    c.Response(w, message, http.StatusBadRequest)
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