package main

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app"
	"github.com/serajam/play-iris/app/handlers"
	"github.com/serajam/play-iris/app/repositories"
)

func main() {
	a := app.Initialize()
	r := repositories.PageRepository{a.Gorm}

	iris.Config().Render.Directory = "app/views"
	iris.Config().Render.Layout = "layout"
	iris.Config().Render.RequirePartials = true

	iris.Handle("GET", "/", handlers.HomePageHandler{&r})
	iris.Handle("GET", "/page/:id", handlers.PageHandler{&r})
	iris.Static("/public", "./public/", 1)

	iris.Listen(":8080")
}
