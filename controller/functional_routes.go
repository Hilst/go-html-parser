package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	"github.com/gin-gonic/gin"
)

const (
	nameQuery           = "name"
	notFoundTmpl        = "404.html"
	mainTmpl            = "MAIN"
	paginationDivStrFmt = "<div id=\"page_%s\" %s>\n%s\n</div>\n"
	hiddenMark          = "hidden"
)

func (c *Controller) get_layout_layoutname(ctx *gin.Context) {
	layoutName := ctx.Query(nameQuery)

	layoutch := make(chan mdl.LayoutResponse)
	go c.service.RequestLayout(layoutName, layoutch)

	datach := make(chan mdl.DataResponse, 1)
	go c.service.RequestData(layoutName, datach)

	var builder strings.Builder
	var hiddenNotation string
	var s string
	i := 0
	for layout := range layoutch {
		if layout.Error != nil {
			ctx.HTML(http.StatusNotFound, notFoundTmpl, layout.Error.Error())
			return
		}
		if i == 0 {
			hiddenNotation = ""
		} else {
			hiddenNotation = hiddenMark
		}
		s = fmt.Sprintf(paginationDivStrFmt, layout.Ok.Name, hiddenNotation, layout.Ok.Tmpl)
		builder.WriteString(s)
		i++
	}
	combinedLayout := builder.String()

	c.ts.ParseLayout(combinedLayout)
	ctx.HTML(http.StatusOK, mainTmpl, <-datach)
}

type TestRequest struct {
	LayoutHTML string `form:"html"`
	Data       string `form:"data"`
}

func (c *Controller) patch_layout_test(ctx *gin.Context) {
	var body TestRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var data map[string]any
	err := json.Unmarshal([]byte(body.Data), &data)
	c.ts.ParseLayout(body.LayoutHTML)
	ctx.HTML(http.StatusOK, mainTmpl, mdl.NewDataResp(data, err))
}
