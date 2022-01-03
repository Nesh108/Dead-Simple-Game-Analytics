package models

import "time"

type Event struct {
	Id			int    `json:"id" gorm:"primaryKey"`
	UserId		string `json:"user_id"`
	ProjectName	string `json:"project_name"`
	EventKey	string `json:"event_key"`
	EventValue	string `json:"event_value"`
	RequestId	string `json:"request_id"`
	Timestamp	time.Time 	`json:"timestamp" gorm:"autoCreateTime"`
}
