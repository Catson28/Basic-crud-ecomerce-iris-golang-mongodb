package main

import (
	//"github.com/iris-contrib/middleware/cors"

	"tentativa/route"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.AllowMethods(iris.MethodOptions)
	app.Logger().SetLevel("debug")

	app.Validator = validator.New()
	return app
}

func main() {
	app := newApp()

	route.InitRoute(app)
	//route.PermissionRoute(app)
	app.Run(iris.Addr(":8080"))
}
