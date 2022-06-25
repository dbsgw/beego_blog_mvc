package ArticleController

import (
	"beego_blog_mvc/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

// Get 渲染创建文章页面
func (c *ArticleController) Get() {
	var results []map[string]interface{}
	err := utils.DB.Raw("select * from sort").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Sort"] = results
	c.TplName = "admin/article/add_article.html"
}

// AddWrite 创建文章
func (c *ArticleController) AddWrite() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort := c.GetString("sort")
	//text = text[0:200] + "......"
	var results []map[string]interface{}
	err := utils.DB.Debug().Raw("insert into article(title, add_time,  sort_id, content_html) VALUES (?,?,?,?)", title, text, utils.GetUnix(), sort, html).Find(&results).Error
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

// GetAll 获取文章列表
func (c *ArticleController) GetAll() {
	var results []map[string]interface{}
	err := utils.DB.Raw("SELECT * FROM article ORDER BY add_time DESC").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("文章获取失败")
	}
	c.Data["Article"] = results
	c.TplName = "admin/article/article.html"
}

// EditWrite 渲染修改文章页面
func (c *ArticleController) EditWrite() {

	id := c.GetString("id")
	result := map[string]interface{}{}
	utils.DB.Raw("select * from article where id = ?", id).Find(&result)
	fmt.Println(result, "id---")
	var results []map[string]interface{}
	err := utils.DB.Debug().Raw("select * from sort").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Sort"] = results

	c.Data["Result"] = result
	fmt.Println(results, result, "=====")
	c.TplName = "admin/article/edit_article.html"
}

// EditWritePost 修改文章
func (c *ArticleController) EditWritePost() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort := c.GetString("sort")
	id := c.GetString("id")
	fmt.Println(title, html, text, sort)
	var results []map[string]interface{}
	err := utils.DB.Debug().Raw("UPDATE beego_mvc_blog.article t SET t.title = ?,content_html=?,t.sort_id=? WHERE t.id = ?", title, html, sort, id).Find(&results).Error
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

// DelWrite 删除文章
func (c *ArticleController) DelWrite() {
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
