package staffs

import (
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
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

// func GetStaffByID(c *gin.Context) {
// 	id := requests.StaffIdRequest{}
// 	if err := c.BindUri(&id); err != nil {
// 		response.BadRequest(c, err.Error())
// 		return
// 	}

// 	data, err := GetByIdStaffService(c, id.ID)
// 	if err != nil {
// 		response.InternalError(c, err.Error())
// 		return
// 	}
// 	response.Success(c, data)
// }

func GetStaffByID(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdStaffService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func StaffList(c *gin.Context) {
	req := requests.StaffRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListStaffService(c.Request.Context(), req)
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

func CreateStaff(c *gin.Context) {
	req := requests.StaffCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateStaffService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Success")
}

func UpdateStaff(c *gin.Context) {
	id := requests.StaffIdRequest{}

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.StaffUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateStaffService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "Admin updated successfully")
}

func DeleteStaff(c *gin.Context) {
	id := requests.StaffIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	err := DeleteStaffService(c, id.ID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "delete successfully")
}
