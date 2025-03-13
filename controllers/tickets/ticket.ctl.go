package tickets

import (
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TicketList(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.TicketRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	data, total, err := ListTicketService(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	paginate := model.Paginate{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}

	response.SuccessWithPaginate(c, data, paginate)
}

func GetTicketByID(c *gin.Context) {
	id := requests.TicketIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdTicketService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

// CreateTicket - API สำหรับสร้างตั๋วและ QR Code
func CreateTicket(c *gin.Context) {
	user := c.GetInt("user_id")
	req := requests.TicketCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	req.UserID = user

	ticket, err := CreateTicketService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// ส่ง Base64 กลับไปที่ Client
	c.JSON(http.StatusOK, gin.H{
		"tick_id": ticket.ID,
		"message": "Ticket created successfully",
		"qr_code": ticket.QrCode,
	})
}
