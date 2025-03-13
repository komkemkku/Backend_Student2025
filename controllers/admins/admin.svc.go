package admin

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
		Where("a.id = ?", id).
		Scan(ctx, admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func CreateAdminService(ctx context.Context, req requests.AdminCreateRequest) (*model.Admins, error) {

	hashpassword, _ := utils.HashPassword(req.Password)

	// เพิ่มadmin
	admin := &model.Admins{
		Username: req.Username,
		Password: hashpassword,
	}
	admin.SetCreatedNow()

	_, err := db.NewInsert().Model(admin).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return admin, nil

}

func UpdateAdminService(ctx context.Context, id int, req requests.AdminUpdateRequest) (*model.Admins, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("id=?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("admin not found")
	}

	admin := &model.Admins{}

	err = db.NewSelect().Model(admin).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	admin.Username = req.Username
	admin.Password = req.Password

	_, err = db.NewUpdate().Model(admin).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func DeleteAdminService(ctx context.Context, id int) error {
	ex, err := db.NewSelect().TableExpr("admins").Where("id=?", id).Exists(ctx)

	if err != nil {
		return err
	}

	if !ex {
		return errors.New("admin not found")
	}

	_, err = db.NewDelete().TableExpr("admins").Where("id =?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
