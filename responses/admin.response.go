package response

type AdminResponses struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}
