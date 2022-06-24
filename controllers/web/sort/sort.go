package sort

import (
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/utils"
	"fmt"
)

type SortController struct {
	web.BaseController
}

func (c *SortController) Get() {
	c.InitData()
	id := c.Ctx.Input.Param(":id")
	result := []map[string]interface{}{}
	utils.DB.Raw("select * from article where sort_id = ?", id).Find(&result)
	c.Data["Arr"] = result
	c.TplName = "web/sort/sort.html"
}
func (c *SortController) Category() {
	c.InitData()
	id := c.Ctx.Input.Param(":id")
	fmt.Println(id)
	result := map[string]interface{}{}
	utils.DB.Raw("select * from article where id = ?", id).Find(&result)
	c.Data["Article"] = result
	c.TplName = "web/category/category.html"
}
