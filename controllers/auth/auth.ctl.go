package auth

import (
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"
	"Beckend_Student2025/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	req := requests.UserLoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginUserService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenUser(c, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func LoginStaff(c *gin.Context) {
	req := requests.StaffAdminLoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginStaffService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenStaff(c, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func LoginAdmin(c *gin.Context) {
	req := requests.StaffAdminLoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := LoginAdminService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	token, err := jwt.GenerateTokenAdmin(c, data)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}
