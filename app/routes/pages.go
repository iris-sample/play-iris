package routes

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app/handlers"
	"github.com/serajam/play-iris/app/repositories"
)

func PagesRoutes(r repositories.PageRepository) {
	iris.Handle("GET", "/", handlers.HomePageHandler{&r})
	iris.Handle("GET", "/page/:id", handlers.PageHandler{&r})
	iris.Static("/public", "./public/", 1)
}
