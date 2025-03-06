package users

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"Beckend_Student2025/utils"
	"context"
	"errors"
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
		Column("u.id", "u.firstname", "u.lastname", "u.nickname", "u.email", "u.password", "u.student_id", "u.faculty", "u.medical_condition", "u.food_allergies", "a.created_at").
		Scan(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUserService(ctx context.Context, req requests.UserCreateRequest) (*model.Users, error) {

	hashpassword, _ := utils.HashPassword(req.Password)

	user := &model.Users{
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		Nickname:         req.Nickname,
		StudentID:        req.StudentID,
		Faculty:          req.Faculty,
		MedicalCondition: req.MedicalCondition,
		FoodAllergies:    req.FoodAllergies,
		Email:            req.Email,
		Password:         hashpassword,
	}
	user.SetCreatedNow()

	_, err := db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil

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
