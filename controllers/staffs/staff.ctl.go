package staffs

import (
	response "Beckend_Student2025/responses"

	"github.com/gin-gonic/gin"
)

func GetInfoStaff(c *gin.Context) {
	admin := c.GetInt("staff_id")

	data, err := GetByIdStaffService(c, admin)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}
