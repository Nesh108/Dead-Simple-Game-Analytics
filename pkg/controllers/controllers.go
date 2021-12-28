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
	w.Header().Add("Content-Type", "application/json")
    var response models.Response
    response.Message = message
    json.NewEncoder(w).Encode(response)
    return
}

func (c controller) SuccessResponse(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
    var response models.Response
    response.Message = "success"
    json.NewEncoder(w).Encode(response)
    return
}