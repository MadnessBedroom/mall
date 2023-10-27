package adminctrl

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsAdminController struct {
	service adminsrv.UmsAdminService
}

type UmsAdminController interface {
	UmsAdminLogin(c *fiber.Ctx) error
	UmsAdminRegister(c *fiber.Ctx) error
}

// UmsAdminLogin 管理员登录
func (uc *umsAdminController) UmsAdminLogin(c *fiber.Ctx) error {
	data := &model.UmsAdminLoginReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}

	data.LoginIpAddr = c.IP()
	data.LoginTime = time.Now()

	if resp, err := uc.service.UmsAdminLogin(c.Context(), data); err == nil {
		return response.Build(c, nil, resp)
	} else {
		return response.Build(c, err, resp)
	}
}

// UmsAdminRegister 管理员注册
func (uc *umsAdminController) UmsAdminRegister(c *fiber.Ctx) error {
	data := &model.UmsAdminRegisterReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}

	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}

	data.RegIpAddr = c.IP()
	if err := uc.service.UmsAdminRegister(c.Context(), data); err != nil {
		return response.Build(c, e.ErrInternalServer().WithMsg("注册失败, 请稍后再试"), nil)
	}

	return response.Build(c, nil, nil)
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
