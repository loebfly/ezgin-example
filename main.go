package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

//@title	x
//@version 	1.0.0
//@description	x
func main() {

	ezgin.Start(app.Start{
		GinCfg: app.GinCfg{
			RecoveryHandler: func(c *gin.Context, err interface{}) {
				c.JSON(http.StatusOK, i18n.SystemError.ErrorRes())
			},
			NoRouteHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, i18n.UrlNotFound.ErrorRes())
			},
			SwaggerRelativePath: "/docs/*any",
			SwaggerHandler:      ginSwagger.WrapHandler(swaggerFiles.Handler),
		},
	})

	ezgin.ShutdownWhenExitSignal()
}
