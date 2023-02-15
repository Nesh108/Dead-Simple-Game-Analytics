package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	query := c.DB

	projectParam, ok := r.URL.Query()["project"]
	if ok && len(projectParam[0]) > 1 {
		project := projectParam[0]
		query = query.Where("project_name = ?", project)
	}

	sinceParam, ok := r.URL.Query()["since"]
	if ok && len(sinceParam[0]) > 0 {
		since, err := time.Parse("02-01-2006", sinceParam[0]) // expected date format is DD-MM-YYYY
		if err != nil {
			c.ValidationErrorResponse(w, "Param since must be formatted DD-MM-YYYY: "+err.Error())
			return
		}
		query = query.Where("timestamp >= ?", since)
	}

	lastIdParam, ok := r.URL.Query()["last_id"]
	if ok && len(lastIdParam[0]) > 0 {
		lastId := lastIdParam[0]
		query = query.Where("id > ?", lastId)
	}

	if result := query.Order("id asc").Find(&events); result.Error != nil {
		c.UnhandledErrorResponse(w, "Failed to fetch events", result.Error)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodeErr := json.NewEncoder(w).Encode(events)
	if encodeErr != nil {
		c.UnhandledErrorResponse(w, "Failed to encode Events Response", encodeErr)
	}
}
