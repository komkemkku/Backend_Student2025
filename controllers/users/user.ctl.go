package users

import (
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
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

func GetUserByID(c *gin.Context) {
	id := requests.UserIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdUserService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func UserList(c *gin.Context) {
	req := requests.UserRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListUserService(c.Request.Context(), req)
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

func CreateUser(c *gin.Context) {
	req := requests.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateUserService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Successfully created")
}

// ForgotStudentIDHandler - API สำหรับกู้คืน Student ID
func ForgotStudentID(c *gin.Context) {
	req := requests.ForgotStudentIDRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request format")
		return
	}

	// เรียกใช้งาน ForgotStudentIDService โดยต้องส่ง `db` เข้าไปด้วย
	err := ForgotStudentIDService(c.Request.Context(), db, req.Email)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "Student ID has been sent to your email")
}


func DeleteUser(c *gin.Context) {
	id := requests.UserIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteUserService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}
