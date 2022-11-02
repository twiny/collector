package api

import "time"

// Data
type Data struct {
	ID           string    `json:"id"`
	Website      string    `json:"website"`
	URL          string    `json:"url"`
	TargetScope  string    `json:"target_scope"`
	ParsePattern string    `json:"parse_pattern"`
	HTMLFile     string    `json:"html_file"`
	CreatedAt    time.Time `json:"created_at"`
	LastVisit    time.Time `json:"last_visite"`
}
