package users

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"Beckend_Student2025/utils"
	"context"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
)

var db = config.Database()

func GetByIdUserService(ctx context.Context, id int) (*response.UserResponses, error) {
	ex, err := db.NewSelect().TableExpr("user").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("user not found")
	}
	user := &response.UserResponses{}

	err = db.NewSelect().TableExpr("users AS u").
		Column("u.id", "u.firstname", "u.lastname", "u.nickname", "u.email", "u.student_id", "u.faculty", "u.medical_condition", "u.food_allergies", "a.created_at").
		Scan(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ListUserService(ctx context.Context, req requests.UserRequest) ([]response.UserResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.UserResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("users AS u").
		Column("u.id", "u.firstname", "u.lastname", "u.nickname", "u.email", "u.student_id", "u.faculty", "u.medical_condition", "u.food_allergies", "a.created_at")

	if req.Search != "" {
		query.Where("u.student_id ILIKE ? OR u.firstname ILIKE ? OR u.lastname ILIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.OrderExpr("u.created_at DESC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func CreateUserService(ctx context.Context, req requests.UserCreateRequest) (*model.Users, error) {

	// hashpassword, _ := utils.HashPassword(req.Password)

	user := &model.Users{
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		Nickname:         req.Nickname,
		StudentID:        req.StudentID,
		Faculty:          req.Faculty,
		MedicalCondition: req.MedicalCondition,
		FoodAllergies:    req.FoodAllergies,
		Email:            req.Email,
		// Password:         hashpassword,
	}
	user.SetCreatedNow()

	_, err := db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil

}

// ForgotStudentIDService ส่ง Student ID ไปยัง Email ของผู้ใช้
func ForgotStudentIDService(ctx context.Context, db *bun.DB, email string) error {
	user := &model.Users{}

	// ค้นหาผู้ใช้จากอีเมล
	err := db.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return errors.New("user not found")
	}

	// สร้างเนื้อหาอีเมล
	subject := "Your Student ID Recovery"
	body := fmt.Sprintf("Hello %s,\n\nYour Student ID is: %s\n\nKhonkaen University,\nStudent Info", user.Firstname, user.StudentID)

	// ส่งอีเมลไปยังผู้ใช้
	err = utils.SendEmail(user.Email, subject, body)
	if err != nil {
		return errors.New("failed to send email")
	}

	return nil
}

func DeleteUserService(ctx context.Context, id int) error {
	ex, err := db.NewSelect().TableExpr("users").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user not found")
	}

	_, err = db.NewDelete().TableExpr("users").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
