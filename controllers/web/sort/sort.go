package SortController

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
)

type SortController struct {
	web.BaseController
}

func (c *SortController) Get() {
	c.InitData()
	id := c.Ctx.Input.Param(":id")
	var article []models.Article
	utils.DB.Where("sort_id = ?", id).Find(&article)
	c.Data["Article"] = article
	c.TplName = "web/sort/sort.html"
}
