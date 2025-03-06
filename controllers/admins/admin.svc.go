package admin

import (
	config "Beckend_Student2025/configs"
	response "Beckend_Student2025/responses"
	"context"
	"errors"
)

var db = config.Database()

func GetByIdAdminService(ctx context.Context, id int) (*response.AdminResponses, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("admin not found")
	}
	admin := &response.AdminResponses{}

	err = db.NewSelect().TableExpr("admins AS a").
		Column("a.id", "a.username", "a.password", "a.created_at").
		Scan(ctx, admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
