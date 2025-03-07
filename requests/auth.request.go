package requests

type UserLoginRequest struct {
	Email     string `json:"email"`
	StudentID string `json:"student_id"`
	// Password  string `json:"password"`
}

type StaffAdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
