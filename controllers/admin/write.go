package admin

import (
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type WriteController struct {
	beego.Controller
}

func (c *WriteController) Get() {
	results := []map[string]interface{}{}
	err := utils.DB.Raw("select * from sort").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Sort"] = results
	c.TplName = "admin/write.html"
}

func (c *WriteController) AddWrite() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort := c.GetString("sort")
	//text = text[0:200] + "......"
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("insert into article(title, main, add_time,  sort_id, content_html) VALUES (?,?,?,?,?)", title, text, utils.GetUnix(), sort, html).Find(&results).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": results,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": results,
		}
	}
	c.ServeJSON() // 返回json

}
