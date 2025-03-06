package admin

import (
	response "Beckend_Student2025/responses"

	"github.com/gin-gonic/gin"
)

func GetInfoAdmin(c *gin.Context) {
	admin := c.GetInt("admin_id")

	data, err := GetByIdAdminService(c, admin)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}
