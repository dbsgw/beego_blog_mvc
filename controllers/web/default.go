package web

import (
	"beego_blog_mvc/utils"
)

type MainController struct {
	BaseController
}

// Get 首页查询接口
func (c *MainController) Get() {
	c.InitData()
	var result []map[string]interface{}
	utils.DB.Raw("select * from article order by add_time desc").Find(&result)

	var results []map[string]interface{}
	utils.DB.Raw("SELECT * FROM `article` ORDER BY RAND() LIMIT 5").Find(&results)
	c.Data["Arr"] = result
	c.Data["Acrid"] = results
	c.TplName = "web/index.html"

}
