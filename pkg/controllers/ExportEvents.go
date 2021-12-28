package controllers

import (
	"encoding/json"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
	"os"
	"strconv"
	"log"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) ExportEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	if result := c.DB.Find(&events); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("{\"message\" :\"Failed to fetch results.\"}")
		return
	}

	filename := time.Now().Format("2006-01-02+15:04:05") + "_export.csv"
    f, err := os.Create(filename)
    defer f.Close()

    if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("{\"message\" :\"Failed to open file.\"}")
		return
    }

    csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()
	csvWriter.Write([]string{"id","user_id","project_name","event_key","event_value","request_id"})
	for _, value := range events {
		log.Println("Writing row: " + strconv.Itoa(value.Id) + ", " + value.UserId + ", " + value.ProjectName + ", " + value.EventKey + ", " + value.EventValue + ", " + value.RequestId)
        err = csvWriter.Write([]string{strconv.Itoa(value.Id),value.UserId,value.ProjectName,value.EventKey,value.EventValue,value.RequestId})
        if err != nil {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("{\"message\" :\"Failed to write value.\"}")
			return
		}
    }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filename)
}
