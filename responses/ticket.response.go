package response

type TicketResponses struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	EventID   int    `json:"event_id"`
	QrCode    string `json:"qr_code"`
	CreatedAt string `json:"created_at"`
}
