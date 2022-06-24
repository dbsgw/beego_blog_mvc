package admin

import (
	"beego_blog_mvc/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	results := []map[string]interface{}{}
	err := utils.DB.Raw("SELECT * FROM article ORDER BY add_time DESC").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("文章获取失败")
	}
	c.Data["Article"] = results
	c.TplName = "admin/admin.html"
}

func (c *AdminController) EditWrite() {

	id := c.GetString("id")
	result := map[string]interface{}{}
	utils.DB.Raw("select * from article where id = ?", id).Find(&result)
	fmt.Println(result, "id---")
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("select * from sort").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Sort"] = results

	c.Data["Result"] = result
	fmt.Println(results, result, "=====")
	c.TplName = "admin/edit_admin.html"
}

func (c *AdminController) EditWritePost() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort := c.GetString("sort")
	id := c.GetString("id")
	fmt.Println(title, html, text, sort)
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("UPDATE beego_mvc_blog.article t SET t.title = ?,t.main=?,content_html=?,t.sort_id=? WHERE t.id = ?", title, text, html, sort, id).Find(&results).Error
	c.Data["Article"] = results
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

// 删除分类
func (c *AdminController) DelWrite() {
	id := c.GetString("id")
	result := map[string]interface{}{}
	err := utils.DB.Raw("delete from article where id = ?", id).Find(&result).Error
	fmt.Println(result, "id---")

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": result,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": result,
		}
	}
	c.ServeJSON() // 返回json
}
