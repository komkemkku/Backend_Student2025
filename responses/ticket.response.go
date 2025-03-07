package response

type TicketResponses struct {
	ID        int       `json:"id"`
	User      UserResp  `json:"user"`
	Event     EventResp `json:"event"`
	QrCode    string    `json:"qr_code"`
	CreatedAt string    `json:"created_at"`
}

type UserResp struct {
	ID               int    `json:"id"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Nickname         string `json:"nickname"`
	Email            string `json:"email"`
	StudentID        string `json:"student_id"`
	Faculty          string `json:"faculty"`
	MedicalCondition string `json:"medical_condition"`
	FoodAllergies    string `json:"food_allergies"`
}

type EventResp struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
