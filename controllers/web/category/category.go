package CategoryController

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
)

type CategoryController struct {
	web.BaseController
}

func (c *CategoryController) Category() {
	c.InitData()
	id := c.Ctx.Input.Param(":id")
	var article models.Article
	utils.DB.Where("id = ?", id).Find(&article)
	c.Data["Article"] = article
	c.TplName = "web/category/category.html"
}
