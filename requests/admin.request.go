package requests

type AdminRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type AdminIdRequest struct {
	ID int `uri:"id"`
}

type AdminCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminUpdateRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
