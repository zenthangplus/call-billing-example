package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zenthangplus/call-billing-example/src/api/controller"
	"github.com/zenthangplus/call-billing-example/src/api/docs"
	"github.com/zenthangplus/call-billing-example/src/api/properties"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App          *golib.App
	Engine       *gin.Engine
	SwaggerProps *properties.SwaggerProperties

	Actuator        *actuator.Endpoint
	OrderController *controller.CallController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	if p.SwaggerProps.Enabled {
		docs.SwaggerInfo.BasePath = p.App.Path()
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	group.POST("/v1/mobile/:username/call", p.OrderController.EndCall)
}
