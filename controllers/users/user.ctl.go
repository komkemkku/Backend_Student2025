package users

import (
	response "Beckend_Student2025/responses"

	"github.com/gin-gonic/gin"
)

func GetInfoUser(c *gin.Context) {
	user := c.GetInt("user_id")

	data, err := GetByIdUserService(c, user)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)

}
