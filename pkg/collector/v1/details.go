package collector

import "time"

// Details
type Details struct {
	ID         string    `json:"id"`
	Website    string    `json:"website"`
	URL        string    `json:"url"`
	PageTitle  string    `json:"page_title"`
	HTMLFile   string    `json:"html_file"`
	FirstVisit time.Time `json:"first_visit"`
	LastVisit  time.Time `json:"last_visit"`
}
