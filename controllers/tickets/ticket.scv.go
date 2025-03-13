package tickets

import (
	config "Beckend_Student2025/configs"
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"Beckend_Student2025/utils"
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var db = config.Database()

func ListTicketService(ctx context.Context, req requests.TicketRequest) ([]response.TicketResponses, int, error) {

	var Offset int64
	if req.Page > 0 {
		Offset = (req.Page - 1) * req.Size
	}

	resp := []response.TicketResponses{}

	// สร้าง query
	query := db.NewSelect().
		TableExpr("tickets AS t").
		Column("t.id", "t.qr_code", "t.created_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.firstname AS user__firstname").
		ColumnExpr("u.lastname AS user__lastname").
		ColumnExpr("u.nickname AS user__nickname").
		ColumnExpr("u.email AS user__email").
		ColumnExpr("u.student_id AS user__student_id").
		ColumnExpr("u.faculty AS user__faculty").
		ColumnExpr("u.medical_condition AS user__medical_condition").
		ColumnExpr("u.food_allergies AS user__food_allergies").
		ColumnExpr("e.id AS event__id").
		ColumnExpr("e.image AS event__image").
		ColumnExpr("e.name AS event__name").
		ColumnExpr("e.description AS event__description").
		ColumnExpr("e.location AS event__location").
		ColumnExpr("e.start_time AS event__start_time").
		ColumnExpr("e.end_time AS event__end_time").
		ColumnExpr("e.start_date AS event__start_date").
		ColumnExpr("e.end_date AS event__end_date").
		Join("LEFT JOIN users AS u ON u.id = t.user_id").
		Join("LEFT JOIN events AS e ON e.id = t.event_id")

	if req.UserID != 0 {
		query.Where("t.user_id = ?", req.UserID)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Execute query
	err = query.Offset(int(Offset)).Limit(int(req.Size)).Scan(ctx, &resp)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func GetByIdTicketService(ctx context.Context, id int) (*response.TicketResponses, error) {
	ex, err := db.NewSelect().TableExpr("tickets").Where("id = ?", id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("ticket not found")
	}
	ticket := &response.TicketResponses{}

	err = db.NewSelect().TableExpr("tickets AS t").
		Column("t.id", "t.qr_code", "t.created_at").
		ColumnExpr("u.id AS user__id").
		ColumnExpr("u.firstname AS user__firstname").
		ColumnExpr("u.lastname AS user__lastname").
		ColumnExpr("u.nickname AS user__nickname").
		ColumnExpr("u.email AS user__email").
		ColumnExpr("u.student_id AS user__student_id").
		ColumnExpr("u.faculty AS user__faculty").
		ColumnExpr("u.medical_condition AS user__medical_condition").
		ColumnExpr("u.food_allergies AS user__food_allergies").
		ColumnExpr("e.id AS event__id").
		ColumnExpr("e.image AS event__image").
		ColumnExpr("e.name AS event__name").
		ColumnExpr("e.description AS event__description").
		ColumnExpr("e.location AS event__location").
		ColumnExpr("e.dress AS event__dress").
		ColumnExpr("e.start_time AS event__start_time").
		ColumnExpr("e.end_time AS event__end_time").
		ColumnExpr("e.start_date AS event__start_date").
		ColumnExpr("e.end_date AS event__end_date").
		Join("LEFT JOIN users AS u ON u.id = t.user_id").
		Join("LEFT JOIN events AS e ON e.id = t.event_id").
		Where("t.id = ?", id).
		Scan(ctx, ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func CreateTicketService(c *gin.Context, req requests.TicketCreateRequest) (*model.Tickets, error) {
	// ดึง user_id จาก Context (Middleware ควรตั้งค่า user_id ไว้ก่อน)
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, fmt.Errorf("user authentication required")
	}

	// ตรวจสอบว่าผู้ใช้ลงทะเบียนแล้วหรือยัง
	var existingTicket model.Tickets
	err := db.NewSelect().Model(&existingTicket).
		Where("user_id = ? AND event_id = ?", userID, req.EventID).
		Scan(c.Request.Context())
	if err == nil {
		return nil, fmt.Errorf("user already registered for this event")
	}

	// สร้างข้อมูลสำหรับ QR Code
	qrData := fmt.Sprintf(`{"user_id": %v, "event_id": %d}`, userID, req.EventID)

	// เรียกใช้งาน utils.GenerateQRCodeBase64
	qrBase64, err := utils.GenerateQRCodeBase64(qrData)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR Code: %v", err)
	}

	// บันทึกตั๋วลงในฐานข้อมูล
	newTicket := &model.Tickets{
		UserID:  req.UserID,
		EventID: req.EventID,
		QrCode:  qrBase64,
	}
	_, err = db.NewInsert().Model(newTicket).Exec(c.Request.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to save ticket: %v", err)
	}

	return newTicket, nil
}
