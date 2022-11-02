package collector

import "time"

// Details
type Details struct {
	ID           string    `json:"id"`
	Website      string    `json:"website"`
	URL          string    `json:"url"`
	Respong      string    `json:"response"`
	TargetScope  string    `json:"target_scope"`
	ParsePattern string    `json:"parse_pattern"`
	HTMLFile     string    `json:"html_file"`
	FirstVisit   time.Time `json:"first_visit"`
	LastVisit    time.Time `json:"last_visit"`
}
