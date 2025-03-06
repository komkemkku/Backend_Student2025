package requests

type StaffRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type StaffIdRequest struct {
	ID int `uri:"id"`
}

type StaffCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StaffUpdateRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
