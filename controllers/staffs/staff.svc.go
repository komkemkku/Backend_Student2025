package staffs

import (
	config "Beckend_Student2025/configs"
	response "Beckend_Student2025/responses"
	"context"
	"errors"
)

var db = config.Database()

func GetByIdStaffService(ctx context.Context, id int) (*response.StaffResponses, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("staff not found")
	}
	staff := &response.StaffResponses{}

	err = db.NewSelect().TableExpr("staffs AS s").
		Column("s.id", "s.username", "s.password", "s.created_at").
		Scan(ctx, staff)
	if err != nil {
		return nil, err
	}
	return staff, nil
}
