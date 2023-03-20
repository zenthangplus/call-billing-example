package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenthangplus/call-billing-example/src/api/resource"
	"github.com/zenthangplus/call-billing-example/src/core/enum"
	"github.com/zenthangplus/call-billing-example/src/core/service"
	"gitlab.com/golibs-starter/golib/web/log"
	"gitlab.com/golibs-starter/golib/web/response"
)

type BillingController struct {
	service service.BillingService
}

func NewBillingController(service service.BillingService) *BillingController {
	return &BillingController{service: service}
}

// GetBilling
// @Summary Get billing for a user
// @Tags BillingController
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Success 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/mobile/{username}/billing [get]
func (s BillingController) GetBilling(c *gin.Context) {
	username := c.Param("username")
	if len(username) == 0 {
		response.WriteError(c.Writer, enum.ErrMissingUsername)
	}
	estBilling, err := s.service.Get(c, username)
	if err != nil {
		log.Error(c, "Cannot get billing for user [%s], err [%s]", username, err)
		response.WriteError(c.Writer, err)
		return
	}
	response.Write(c.Writer, response.Ok(resource.NewBillingResponse(estBilling)))
}
