package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenthangplus/call-billing-example/src/api/resource"
	"github.com/zenthangplus/call-billing-example/src/core/enum"
	"github.com/zenthangplus/call-billing-example/src/core/service"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/golibs-starter/golib/web/response"
)

type CallController struct {
	service service.CallService
}

func NewCallController(service service.CallService) *CallController {
	return &CallController{service: service}
}

// EndCall
// @Summary End a call
// @Tags CallController
// @Accept  json
// @Produce  json
// @Param	username	path	string	true	"username"
// @Param request body resource.EndCallRequest true "Request body"
// @Success 200 {object} response.Response
// @Success 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/mobile/{username}/call [put]
func (s CallController) EndCall(c *gin.Context) {
	username := c.Param("username")
	if len(username) == 0 {
		response.WriteError(c.Writer, enum.ErrMissingUsername)
	}
	var body resource.EndCallRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Warn(c, "Cannot bind request body, err [%s]", err)
		response.WriteError(c.Writer, baseEx.BadRequest)
		return
	}
	_, err := s.service.EndCall(c, username, body.CallDuration)
	if err != nil {
		log.Error(c, "Cannot end a call, err [%s]", err)
		response.WriteError(c.Writer, err)
		return
	}
	response.Write(c.Writer, response.Ok(nil))
}
