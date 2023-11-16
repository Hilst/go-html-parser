package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) get_tab_sample(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "sample.html", nil)
}

func (c *Controller) get_tab_test(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "tester.html", nil)
}
