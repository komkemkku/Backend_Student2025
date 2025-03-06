package response

type CheckInResponses struct {
	ID        int   `json:"id"`
	TicketID  int   `json:"ticket_id"`
	StaffID   int   `json:"staff_id"`
	CreatedAt int64 `json:"created_at"`
}
