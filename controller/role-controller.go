package controller

import (
	"net/http"

	"github.com/Lasang3012/facilityManage/entity"
	"github.com/Lasang3012/facilityManage/helper"
	"github.com/Lasang3012/facilityManage/service"
	"github.com/gin-gonic/gin"
)

//RoleController is a ...
type RoleController interface {
	All(context *gin.Context)
}

type roleController struct {
	roleService service.RoleService
}

//NewRoleController create a new instances of BoookController
func NewRoleController(roleServ service.RoleService) RoleController {
	return &roleController{
		roleService: roleServ,
	}
}

func (c *roleController) All(context *gin.Context) {
	var books []entity.Role = c.roleService.All()
	res := helper.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, res)
}
