package main

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app"
	"github.com/serajam/play-iris/app/handlers"
	"github.com/serajam/play-iris/app/repositories"
)

func main() {
	a := app.Initialize()
	r := repositories.PageRepository{a.Db}

	iris.Config().Render.Directory = "app/views"
	iris.Handle("GET", "/", handlers.HomePageHandler{&r})
	iris.Handle("GET", "/page/:id", handlers.PageHandler{&r})
	iris.Listen(":8080")
}
