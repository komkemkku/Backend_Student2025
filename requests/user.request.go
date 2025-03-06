package requests

type UserRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type UserIdRequest struct {
	ID int `uri:"id"`
}

type UserCreateRequest struct {
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Nickname         string `json:"nickname"`
	StudentID        string `json:"student_id"`
	Faculty          string `json:"faculty"`
	MedicalCondition string `json:"medical_condition"`
	FoodAllergies    string `json:"food_allergies"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}
