package response

type UserResponses struct {
	ID               int    `json:"id"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Nickname         string `json:"nickname"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	StudentID        string `json:"student_id"`
	Faculty          string `json:"faculty"`
	MedicalCondition string `json:"medical_condition"`
	FoodAllergies    string `json:"food_allergies"`
	CreatedAt        int64    `json:"created_at"`
}
