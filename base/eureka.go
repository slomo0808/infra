package base

import (
	"github.com/kataras/iris"
	"github.com/slomo0808/infra"
	"github.com/tietang/go-eureka-client/eureka"
	"time"
)

type EurekaStarter struct {
	infra.BaseStarter
	client *eureka.Client
}

func (s *EurekaStarter) Init(ctx infra.StarterContext) {
	s.client = eureka.NewClient(ctx.Props())
	s.client.Start()
}

func (s *EurekaStarter) Setup(ctx infra.StarterContext) {
	info := make(map[string]interface{})
	info["startTime"] = time.Now()
	info["appName"] = ctx.Props().GetDefault("app.name", "resk")
	Iris().Get("/info", func(ctx iris.Context) {
		ctx.JSON(info)
	})

	Iris().Get("/health", func(ctx iris.Context) {
		health := eureka.Health{
			Details: make(map[string]interface{}),
		}
		health.Status = eureka.StatusUp

		ctx.JSON(health)
	})
}
