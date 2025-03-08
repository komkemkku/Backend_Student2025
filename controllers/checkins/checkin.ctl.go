package checkins

import (
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckinCreate(c *gin.Context) {
	// ดึง Staff ID จาก Token (Middleware ควรใส่ user_id ไว้)
	staffID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "Unauthorized: Missing staff ID")
		return
	}

	req := requests.CheckInCreateRequest{}

	// ดึง JSON ข้อมูลจาก QR Code
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid QR Code format")
		return
	}

	// ใส่ StaffID ลงไปใน Request
	req.StaffID, _ = staffID.(int)

	// เรียกใช้ Checkin Service
	checkinData, err := CheckinService(c, req)
	if err != nil {
		if err.Error() == "user already checked in" {
			response.BadRequest(c, "User has already checked in")
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	// ส่งข้อมูล User และ Event กลับไปที่ Frontend
	c.JSON(http.StatusOK, gin.H{
		"message":    "Check-in successful",
		"user_info":  checkinData.User,
		"event_info": checkinData.Event,
		"checked_in": checkinData.CheckedInAt,
	})
}

