package handlers

import (
	"github.com/kataras/iris"
	"github.com/serajam/play-iris/app/repositories"
	"fmt"
	"strconv"
	"log"
)

const param_id = "id"
const page_template = "pages/page"

type PageHandler struct {
	Repository *repositories.PageRepository
}

func (h PageHandler) Serve(c *iris.Context) {
	p := h.Repository.GetPage(getIdParam(c))

	if p.Id == 0 {
		c.EmitError(404)
	} else {
		c.Render(page_template, p)
	}
}

func getIdParam(c *iris.Context) int {
	id, err := strconv.Atoi(c.Param(param_id))
	errCheck(c.Param(param_id), err)
	return id
}

func errCheck(param string, err error) {
	if err != nil {
		log.Printf(fmt.Sprintf("Invalid parameter for pageId: %s", param))
	}
}
