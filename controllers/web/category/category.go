package category

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/utils"
	"fmt"
)

type CategoryController struct {
	web.BaseController
}

func (c *CategoryController) Category() {
	c.InitData()
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)
	result := map[string]interface{}{}
	utils.DB.Raw("select * from article where id = ?", id).Find(&result)
	c.Data["Article"] = result
	c.TplName = "web/category/category.html"
}
