package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("{\"message\" :\"Failed to read body.\"}")
		return
	}

	var event models.Event
	if errUnmashaling := json.Unmarshal(body, &event); errUnmashaling != nil {
		json.NewEncoder(w).Encode("{\"message\" :\"Failed to parse arguments.\"}")
		return
	}

	if result := c.DB.Create(&event); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
