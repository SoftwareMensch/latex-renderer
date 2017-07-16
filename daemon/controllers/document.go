package controllers

import (
	"github.com/astaxie/beego"
	daemonDto "rkl.io/latex-renderer/daemon/dto"
	"fmt"
	"rkl.io/latex-renderer/contract"
	"rkl.io/latex-renderer/di"
	"time"
	"bytes"
	"io"
)

type DocumentController struct {
	beego.Controller
}

func (c *DocumentController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *DocumentController) ServeJsonException(code int16, message string) {
	c.Ctx.Output.SetStatus(int(code))
	c.Data["json"] = daemonDto.NewException(code, message)
	c.ServeJSON()
	c.Abort(string(code))
}
// Post ...
// @Title Create
// @Description create Foobar
// @Param	body		body 	[]byte	true		"body of tex document"
// @Success 201 {object} []byte
// @Failure 403 body is empty
// @router / [post]
func (c *DocumentController) Post() {
	file, _, err := c.Ctx.Request.FormFile("document")
	if err != nil {
		c.ServeJsonException(400, "no document file found")
	}
	defer file.Close()

	var buffer bytes.Buffer

	io.Copy(&buffer, file)

	service, err := di.GetContainer().GetService("app.renderer")
	if err != nil {
		c.ServeJsonException(500, "no renderer configured")
	}
	renderer := service.(contract.RendererInterface)

	pdfData, err := renderer.Render(buffer.Bytes())
	if err != nil {
		c.ServeJsonException(500, "renderer error")
	}

	t := time.Now()

	c.Ctx.Output.ContentType("application/pdf")
	c.Ctx.Output.Header(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%s-document.pdf\"", t.Format("2006-01-02 15:04:05")),
	)
	c.Ctx.Output.Body(pdfData)
}
