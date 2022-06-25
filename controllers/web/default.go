package web

import (
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
)

type MainController struct {
	BaseController
}

// Get 首页查询接口
func (c *MainController) Get() {
	c.InitData()
	var article []models.Article
	utils.DB.Order("created_at desc").Find(&article)
	var Acrid []models.Article
	utils.DB.Order("RAND()").Limit(5).Find(&Acrid)
	c.Data["Article"] = article
	c.Data["Acrid"] = Acrid
	c.TplName = "web/index.html"
}
