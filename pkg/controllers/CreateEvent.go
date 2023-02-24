package controllers

import (
	"net/http"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
	uuid "github.com/nu7hatch/gouuid"
)

func (c controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projectParam, ok := r.URL.Query()["project"]
	if !ok || len(projectParam[0]) < 1 {
		c.ValidationErrorResponse(w, "Url param project is required.")
		return
	}
	project := projectParam[0]

	userId := r.PostForm.Get("user_id")
	if userId == "" {
		c.ValidationErrorResponse(w, "Post field user_id is required.")
		return
	}

	eventKeys := r.PostForm["event_key"]
	if len(eventKeys) <= 0 {
		c.ValidationErrorResponse(w, "Post field event_key is required.")
		return
	}

	eventValues := r.PostForm["event_value"]
	if len(eventValues) <= 0 {
		c.ValidationErrorResponse(w, "Post field event_value is required.")
		return
	}

	eventNum := len(eventValues)
	if eventNum != len(eventKeys) {
		c.ValidationErrorResponse(w, "The number of event_value does not match the number of event_key.")
		return
	}

	requestId, err := uuid.NewV4()
	if err != nil {
		c.ValidationErrorResponse(w, "Failed to generate request_id.")
		return
	}

	events := []models.Event{}
	for i := 0; i < eventNum; i++ {
		e := models.Event{
			ProjectName: project,
			UserId:      userId,
			EventKey:    eventKeys[i],
			EventValue:  eventValues[i],
			RequestId:   requestId.String(),
		}
		events = append(events, e)
	}

	tx := c.DB.Begin()
	if result := tx.Create(&events); result.Error != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()
		c.UnhandledErrorResponse(w, "Failed to create event", result.Error)
		return
	}
	// Commit the transaction if everything is successful
	tx.Commit()

	c.SuccessResponse(w)
}
