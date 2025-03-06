package response

type EventResponses struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Is_active   bool   `json:"is_active"`
	CreatedAt   int64  `json:"created_at"`
}
