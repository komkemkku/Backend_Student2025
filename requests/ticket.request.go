package requests

type TicketRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type TicketIdRequest struct {
	ID int `uri:"id"`
}

type TicketCreateRequest struct {
	UserID  int    `json:"user_id"`
	EventID int    `json:"event_id"`
	QrCode  string `json:"qr_code"`
}
