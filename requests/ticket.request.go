package requests

type TicketRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
	UserID int    `form:"user_id"`
}

type TicketIdRequest struct {
	ID int `uri:"id"`
}

type TicketCreateRequest struct {
	UserID  int    `json:"user_id"`
	EventID int    `json:"event_id"`
}
