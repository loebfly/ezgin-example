package main

import (
	"ezgin-example/router"
	"github.com/loebfly/ezgin"
)

func main() {

	// 不带参数的初始化
	ezgin.Start()

	// 带参数的初始化
	//ezgin.Start(app.Start{
	//	YmlPath: "~/ezgin-example/ezgin-example.yml", // yml配置文件的绝对路径
	//	GinCfg: app.GinCfg{
	//		Engine:              nil, // gin的engine实例, 如果为nil, 则会自动创建gin的engine实例
	//		RecoveryHandler:     nil, // gin的RecoveryHandler, 如果为nil, 则会使用默认的RecoveryHandler
	//		NoRouteHandler:      nil, // gin的NoRouteHandler, 如果为nil, 则会使用默认的NoRouteHandler
	//		SwaggerRelativePath: "",  // swagger的相对路径, 如果为"", 则不会启动swagger
	//		SwaggerHandler:      nil, // swagger的handler, 如果为nil, 则会使用默认的swagger handler
	//	},
	//})
	ezgin.Logs.Debug("ezgin 初始化成功")

	// 业务代码
	// ...
	ezgin.Logs.CDebug("ROUTER", "开始配置路由")
	router.Setup()

	// 不带参数的监听退出信号优雅退出服务
	ezgin.ShutdownWhenExitSignal()

	// 带参数的监听退出信号优雅退出服务
	//ezgin.ShutdownWhenExitSignal(app.Shutdown{
	//	WillHandler: nil, // 退出前的回调函数, 如果为nil, 则不会执行回调函数
	//	DidHandler:  nil, // 退出后的回调函数, 如果为nil, 则不会执行回调函数
	//})
}
