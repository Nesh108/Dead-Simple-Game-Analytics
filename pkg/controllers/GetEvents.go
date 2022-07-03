package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	projectParam, ok := r.URL.Query()["project"]
	if !ok || len(projectParam[0]) < 1 {
		if result := c.DB.Order("id asc").Find(&events); result.Error != nil {
			c.UnhandledErrorResponse(w, result.Error)
			return
		}
	} else {
		project := projectParam[0]
		if result := c.DB.Where("project_name = ?", project).Order("id asc").Find(&events); result.Error != nil {
			c.UnhandledErrorResponse(w, result.Error)
			return
		}

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
