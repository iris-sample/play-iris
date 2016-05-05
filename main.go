package main

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app"
	"os"
	"fmt"
)

func main() {

	app.Initialize()

	iris.Config().Render.Gzip = true
	iris.Config().Render.Directory = "app/views"
	iris.Config().Render.Layout = "layout"
	iris.Config().Render.RequirePartials = true

	fmt.Printf("Listen on port %s", os.Getenv("LISTEN_PORT"))
	iris.Listen(os.Getenv("LISTEN_PORT"))
}
