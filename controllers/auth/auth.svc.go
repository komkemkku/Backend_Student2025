package auth

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	"Beckend_Student2025/utils"
	"context"
	"errors"
)

var db = config.Database()

func LoginUserService(ctx context.Context, req requests.UserLoginRequest) (*model.Users, error) {
	ex, err := db.NewSelect().TableExpr("users").Where("email = ?", req.Email).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("email or password not found")
	}

	user := &model.Users{}

	err = db.NewSelect().Model(user).Where("email =?", req.Email).Scan(ctx)
	if err != nil {
		return nil, err
	}

		// ตรวจสอบ student_id แทน password
		if req.StudentID != user.StudentID {
			return nil, errors.New("email or student ID not found")
		}

	// bool := utils.CheckPasswordHash(req.Password, user.Password)

	// if !bool {
	// 	return nil, errors.New("email or password not found")
	// }

	return user, nil
}

func LoginAdminService(ctx context.Context, req requests.StaffAdminLoginRequest) (*model.Admins, error) {
	ex, err := db.NewSelect().TableExpr("admins").Where("username = ?", req.Username).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("username or password not found")
	}

	admin := &model.Admins{}

	err = db.NewSelect().Model(admin).Where("username =?", req.Username).Scan(ctx)
	if err != nil {
		return nil, err
	}

	bool := utils.CheckPasswordHash(req.Password, admin.Password)

	if !bool {
		return nil, errors.New("username or password not found")
	}

	return admin, nil
}

func LoginStaffService(ctx context.Context, req requests.StaffAdminLoginRequest) (*model.Staffs, error) {
	ex, err := db.NewSelect().TableExpr("staffs").Where("username = ?", req.Username).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("username or password not found")
	}

	staff := &model.Staffs{}

	err = db.NewSelect().Model(staff).Where("username =?", req.Username).Scan(ctx)
	if err != nil {
		return nil, err
	}

	bool := utils.CheckPasswordHash(req.Password, staff.Password)

	if !bool {
		return nil, errors.New("username or password not found")
	}

	return staff, nil
}
