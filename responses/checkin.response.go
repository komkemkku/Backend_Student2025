package response

import model "Beckend_Student2025/models"

type CheckInResponse struct {
	ID          int          `json:"id"`
	User        model.Users  `json:"user"`
	Event       model.Events `json:"event"`
	CheckedInAt int64        `json:"checked_in_at"`
}
