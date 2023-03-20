package bootstrap

import (
	"github.com/zenthangplus/call-billing-example/src/adapter/publisher"
	"github.com/zenthangplus/call-billing-example/src/adapter/repository/mysql"
	"github.com/zenthangplus/call-billing-example/src/api/controller"
	"github.com/zenthangplus/call-billing-example/src/api/listener"
	"github.com/zenthangplus/call-billing-example/src/api/properties"
	"github.com/zenthangplus/call-billing-example/src/api/router"
	"github.com/zenthangplus/call-billing-example/src/core/port"
	"github.com/zenthangplus/call-billing-example/src/core/service"
	"gitlab.com/golibs-starter/golib"
	golibdata "gitlab.com/golibs-starter/golib-data"
	golibgin "gitlab.com/golibs-starter/golib-gin"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golib.HttpRequestLogOpt(),

		// Provide datasource auto config
		golibdata.DatasourceOpt(),

		// Provide all application properties
		golib.ProvideProps(properties.NewSwaggerProperties),
		golib.ProvideProps(properties.NewCallProperties),
		golib.ProvideProps(properties.NewBillingProperties),

		// Provide port's implements
		fx.Provide(fx.Annotate(
			publisher.NewEventPublisherAdapter,
			fx.As(new(port.EventPublisher)),
		)),
		fx.Provide(fx.Annotate(
			mysql.NewCallRepositoryAdapter,
			fx.As(new(port.CallRepository)),
		)),
		fx.Provide(fx.Annotate(
			mysql.NewBillingRepositoryAdapter,
			fx.As(new(port.BillingRepository)),
		)),

		// Provide service's configs
		fx.Provide(NewCallConfig),
		fx.Provide(NewBillingConfig),

		// Provide services
		fx.Provide(fx.Annotate(
			service.NewDefaultCallService,
			fx.As(new(service.CallService)),
		)),
		fx.Provide(fx.Annotate(
			service.NewDefaultBillingService,
			fx.As(new(service.BillingService)),
		)),

		// Provide controllers, these controllers will be used
		// when register router was invoked
		fx.Provide(controller.NewCallController),
		fx.Provide(controller.NewBillingController),

		// Provide application listeners
		golib.ProvideEventListener(listener.NewBillingAggregation),

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
	)
}
