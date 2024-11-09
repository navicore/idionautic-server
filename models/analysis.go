package models

type AnalysisResponse struct {
	StartTime      string         `json:"startTime"`
	EndTime        string         `json:"endTime"`
	EventTypeStats map[string]int `json:"eventTypeStats"`
	TargetStats    map[string]int `json:"targetStats"`
}
