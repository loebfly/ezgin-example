package router

import (
	"ezgin-example/controller"
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/engine"
)

func Setup() {
	// 中间件配置
	ezgin.Engine.Use()
	// 路由配置
	ezgin.Engine.Get("user/info", controller.User.Login)
	ezgin.Engine.Post("user/login", controller.User.Login)
	// ...

	// 批量路由配置
	ezgin.Engine.Routers(engine.Put, map[string]engine.HandlerFunc{
		"user/update": controller.User.Login,
	})

	// 批量路由配置
	ezgin.Engine.FreeRouters(map[engine.HttpMethod]map[string]engine.HandlerFunc{
		engine.Any: {
			"user/register": controller.User.Login,
		},
	})

	// 批量分组路由配置
	ezgin.Engine.Group("user").Routers(engine.Put, map[string]engine.HandlerFunc{
		"query": controller.User.Login,
	})

	// 批量分组路由配置
	ezgin.Engine.Group("user").FreeRouters(map[engine.HttpMethod]map[string]engine.HandlerFunc{
		engine.Any: {
			"like": controller.User.Login,
		},
	})
}
