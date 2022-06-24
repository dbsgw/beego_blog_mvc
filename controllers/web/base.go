package web

import (
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

/*
首页查询接口
*/
func (c *BaseController) InitData() {

	results := []map[string]interface{}{}
	utils.DB.Raw("select * from sort").Find(&results)

	c.Data["Sort"] = results
	cookie := c.Ctx.Input.Cookie("token")

	c.Data["Cookie"] = cookie

}
