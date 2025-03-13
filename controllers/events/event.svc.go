package events

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"context"
	"errors"
)

var db = config.Database()

func ListEventService(ctx context.Context, req requests.EventRequest) ([]response.EventResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.EventResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("events AS e").
		Column("e.id", "e.image", "e.name", "e.description", "e.location", "e.dress", "e.start_time", "e.end_time", "e.start_date", "e.end_date", "e.is_active", "e.created_at")

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.Order("e.id ASC").Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdEventService(ctx context.Context, id int) (*response.EventResponses, error) {
	ex, err := db.NewSelect().TableExpr("events").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("event not found")
	}
	event := &response.EventResponses{}

	err = db.NewSelect().TableExpr("events AS e").
		Column("e.id", "e.image", "e.name", "e.description", "e.location", "e.dress", "e.start_time", "e.end_time", "e.start_date", "e.end_date", "e.is_active", "e.created_at").
		Where("e.id = ?", id).
		Scan(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func CreateEventService(ctx context.Context, req requests.EventCreateRequest) (*model.Events, error) {

	// ตรวจสอบชื่อซ้ำ
	exists, err := db.NewSelect().
		TableExpr("events").
		Where("name = ?", req.Name).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("event already exists")
	}

	// เพิ่ม
	event := &model.Events{
		Image:       req.Image,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Dress:       req.Dress,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Is_active:   req.Is_active,
	}
	event.SetCreatedNow()

	_, err = db.NewInsert().Model(event).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return event, nil

}

func UpdateEventService(ctx context.Context, ID int, req requests.EventUpdateRequest) (*model.Events, error) {

	exists, err := db.NewSelect().TableExpr("events").Where("id=?", ID).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("event not found")
	}

	event := &model.Events{}
	err = db.NewSelect().Model(event).Where("id = ?", ID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	// เพิ่ม
	event.Image = req.Image
	event.Name = req.Name
	event.Description = req.Description
	event.Location = req.Location
	event.Dress = req.Dress
	event.StartTime = req.StartTime
	event.EndTime = req.EndTime
	event.StartDate = req.StartDate
	event.EndDate = req.EndDate
	event.Is_active = req.Is_active
	event.SetCreatedNow()

	_, err = db.NewUpdate().Model(event).Where("id =?", ID).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return event, nil

}
