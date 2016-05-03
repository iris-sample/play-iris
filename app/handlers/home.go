package handlers

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app/repositories"
)

const home_page_template = "pages/home"

type HomePageHandler struct {
	Repository *repositories.PageRepository
}

func (h HomePageHandler) Serve(c *iris.Context) {
	p := h.Repository.GetPage(1)

	if p.Id == 0 {
		c.EmitError(404)
	} else {
		c.Render(home_page_template, p)
	}
}
