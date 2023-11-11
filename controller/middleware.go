package controller

// import (
// 	"bytes"

// 	"github.com/gin-gonic/gin"
// )

// func (c *Controller) ginMinifyHTML() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		wb := &bodyWriter{
// 			body:           &bytes.Buffer{},
// 			ResponseWriter: ctx.Writer,
// 		}
// 		ctx.Writer = wb
// 		ctx.Next()
// 		data := wb.body

// 		resp := bytes.NewBuffer([]byte{})
// 		err := c.m.Minify(ctx.ContentType(), resp, data)

// 		wb.ResponseWriter.Write(resp.Bytes())
// 		if err != nil {
// 			wb.body.Reset()
// 		}
// 	}
// }

// type bodyWriter struct {
// 	gin.ResponseWriter
// 	body *bytes.Buffer
// }

// func (w bodyWriter) Write(b []byte) (int, error) {
// 	return w.body.Write(b)
// }
