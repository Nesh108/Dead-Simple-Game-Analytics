package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) GetEvents(w http.ResponseWriter, r *http.Request) {
    var events []models.Event

    if result := c.DB.Find(&events); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(events)
}
