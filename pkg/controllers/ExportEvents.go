package controllers

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"
	"os/exec"
	"fmt"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
)

func (c controller) ExportEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	exportName := "complete"
	projectParam, ok := r.URL.Query()["project"]
    if !ok || len(projectParam[0]) < 1 {
		if result := c.DB.Find(&events); result.Error != nil {
			c.UnhandledErrorResponse(w, result.Error)
			return
		}
    } else {
		project := projectParam[0]
		exportName = project
		if result := c.DB.Where("project_name = ?", project).Find(&events); result.Error != nil {
			c.UnhandledErrorResponse(w, result.Error)
			return
		}
	}

	filename := "exports/" + exportName + "_export.csv"
    f, err := os.Create(filename)
    defer f.Close()

    if err != nil {
		c.UnhandledErrorResponse(w, err)
		return
    }

    writer := csv.NewWriter(f)
	writer.Write([]string{"id","user_id","project_name","event_key","event_value","request_id","timestamp"})

	for _, value := range events {
		err = writer.Write([]string{
			strconv.Itoa(value.Id),
			value.UserId,
			value.ProjectName,
			value.EventKey,
			value.EventValue,
			value.RequestId,
			value.Timestamp.Format("2006-01-02 15:04:05")})
        
		if err != nil {
			writer.Flush()
			c.UnhandledErrorResponse(w, err)
			return
		}
    }
	writer.Flush()
	output, errOutput := exec.Command(os.Getenv("EXPORT_COMMAND_PATH")).Output()
	fmt.Printf("OUT: %s\n", output)
    if errOutput != nil {
        c.UnhandledErrorResponse(w, err)
		return
    }

	c.SuccessResponse(w)
	return
}
