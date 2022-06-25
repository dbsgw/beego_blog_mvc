package web

import (
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

// InitData 首页查询接口
func (c *BaseController) InitData() {
	var results []models.Sort
	utils.DB.Find(&results)
	c.Data["Sort"] = results
	cookie := c.Ctx.Input.Cookie("token")
	c.Data["Cookie"] = cookie
}
