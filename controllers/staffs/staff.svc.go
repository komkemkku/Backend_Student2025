package staffs

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

func CreateStaffService(ctx context.Context, req requests.StaffCreateRequest) (*model.Staffs, error) {

	hashpassword, _ := utils.HashPassword(req.Password)

	staff := &model.Staffs{
		Username: req.Username,
		Password: hashpassword,
	}
	staff.SetCreatedNow()

	_, err := db.NewInsert().Model(staff).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return staff, nil

}

func UpdateStaffService(ctx context.Context, id int, req requests.StaffUpdateRequest) (*model.Staffs, error) {
	ex, err := db.NewSelect().TableExpr("staffs").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("staff not found")
	}

	staff := &model.Staffs{}

	err = db.NewSelect().Model(staff).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	staff.Username = req.Username
	staff.Password = req.Password

	_, err = db.NewUpdate().Model(staff).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func DeleteStaffService(ctx context.Context, id int) error {
	ex, err := db.NewSelect().TableExpr("staffs").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("staff not found")
	}

	_, err = db.NewDelete().TableExpr("staffs").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
