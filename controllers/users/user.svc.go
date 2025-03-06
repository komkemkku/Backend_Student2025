package users

import (
	config "Beckend_Student2025/configs"
	response "Beckend_Student2025/responses"
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
