package handlers

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app/repositories"
	"github.com/serajam/play-iris/app/models"
)

const home_page_template = "pages/home"

type PagesCollection struct {
	Pages []*models.Page
}

type HomePageHandler struct {
	Repository *repositories.PageRepository
}

func (h HomePageHandler) Serve(c *iris.Context) {
	p := h.Repository.GetPagesWithLimit(10)
	pages := PagesCollection{p}

	c.Render(home_page_template, pages)
}
