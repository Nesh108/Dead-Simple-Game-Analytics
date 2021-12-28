package controllers

import (
	"net/http"

	"github.com/nu7hatch/gouuid"
	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projectName, ok := r.URL.Query()["project_name"]
    if !ok || len(projectName[0]) < 1 {
        c.ErrorResponse(w, "Url param project_name is required.")
		return
    }

	userId := r.PostForm.Get("user_id")
	if userId == "" {
		c.ErrorResponse(w, "Post field user_id is required.")
		return
	}

	eventKey := r.PostForm.Get("event_key")
	if eventKey == "" {
		c.ErrorResponse(w, "Post field event_key is required.")
		return
	}

	eventValue := r.PostForm.Get("event_value")
	if eventValue == "" {
		c.ErrorResponse(w, "Post field event_value is required.")
		return
	}
	
	requestId, err := uuid.NewV4()
	if err != nil {
        c.ErrorResponse(w, "Failed to generate request_id.")
		return
    }

	var event models.Event
	event.ProjectName = projectName[0]
	event.UserId = userId
	event.EventKey = eventKey
	event.EventValue = eventValue
	event.RequestId = requestId.String()
	
	if result := c.DB.Create(&event); result.Error != nil {
		c.ErrorResponse(w, "Failed to create event in DB.")
		return
	}

	c.SuccessResponse(w)
}
