package requests

type CheckInRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CheckInIdRequest struct {
	ID int `uri:"id"`
}

type CheckInCreateRequest struct {
	TicketID int `json:"ticket_id"`
	StaffID  int `json:"staff_id"`
}

type CheckInUpdateRequest struct {
	ID       int `json:"id"`
	TicketID int `json:"ticket_id"`
	StaffID  int `json:"staff_id"`
}
