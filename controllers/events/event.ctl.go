package events

import (
	model "Beckend_Student2025/models"
	"Beckend_Student2025/requests"
	response "Beckend_Student2025/responses"

	"github.com/gin-gonic/gin"
)

func GetEventByID(c *gin.Context) {
	id := requests.EventIdRequest{}
	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, err := GetByIdEventService(c, int(id.ID))
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, data)
}

func EventList(c *gin.Context) {
	req := requests.EventRequest{}
	if err := c.BindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	data, total, err := ListEventService(c.Request.Context(), req)
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

func CreateEvent(c *gin.Context) {
	req := requests.EventCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := CreateEventService(c, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, "Successfully created")
}

func UpdateEvent(c *gin.Context) {
	id := requests.EventIdRequest{}

	if err := c.BindUri(&id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	req := requests.EventUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, err := UpdateEventService(c, id.ID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, "event updated successfully")
}