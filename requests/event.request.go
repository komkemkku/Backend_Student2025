package requests

type EventRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type EventIdRequest struct {
	ID int `uri:"id"`
}

type EventCreateRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Dress       string `json:"dress"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Is_active   bool   `json:"is_active"`
}

type EventUpdateRequest struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Dress       string `json:"dress"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Is_active   bool   `json:"is_active"`
}
