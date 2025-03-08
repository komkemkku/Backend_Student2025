package checkins

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"context"
	"fmt"
	"time"
)

var db = config.Database()

// CheckinService - บันทึก Check-in และดึงข้อมูลผู้ใช้ + กิจกรรม
func CheckinService(c context.Context, req requests.CheckInCreateRequest) (*response.CheckInResponse, error) {

	var ticket model.Tickets
	err := db.NewSelect().Model(&ticket).
		Where("user_id = ? AND event_id = ?", req.UserID, req.EventID).
		Scan(c)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket or user not registered")
	}

	// ตรวจสอบว่าผู้ใช้เคย Check-in แล้วหรือไม่
	var existingCheckin model.Checkins
	err = db.NewSelect().Model(&existingCheckin).
		Where("ticket_id = ?", ticket.ID).
		Scan(c)
	if err == nil {
		return nil, fmt.Errorf("user already checked in")
	}


	newCheckin := &model.Checkins{
		TicketID:    ticket.ID,
		StaffID:     req.StaffID, // Staff ID ได้จาก Token
		CheckedInAt: time.Now().Unix(),
	}
	_, err = db.NewInsert().Model(newCheckin).Exec(c)
	if err != nil {
		return nil, fmt.Errorf("failed to check in")
	}

	var user model.Users
	var event model.Events

	err = db.NewSelect().Model(&user).Where("id = ?", req.UserID).Scan(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data")
	}

	err = db.NewSelect().Model(&event).Where("id = ?", req.EventID).Scan(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get event data")
	}

	// ส่งข้อมูลกลับไป
	return &response.CheckInResponse{
		ID:          newCheckin.ID,
		User:        user,
		Event:       event,
		CheckedInAt: newCheckin.CheckedInAt,
	}, nil
}
