package models

import "time"

type PeopleTask struct {
	IdName    int           `json:"id_name"`
	IdTask    int           `json:"id_task"`
	StartTime time.Time     `json:"start_time"`
	StopTime  time.Time     `json:"stop_time"`
	Duration  time.Duration `json:"duration"`
}
