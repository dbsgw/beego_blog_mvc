package ArticleController

import (
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// Get 渲染创建文章页面
func (c *ArticleController) Get() {
	var sort []models.Sort
	err := utils.DB.Find(&sort).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Sort"] = sort
	c.TplName = "admin/article/add_article.html"
}

// AddWrite 创建文章
func (c *ArticleController) AddWrite() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort, err := c.GetInt("sort")
	if err != nil {
		c.Ctx.WriteString("分类id错误")
		return
	}
	//text = text[0:200] + "......"
	article := models.Article{
		Model: models.Model{
			CreatedAt: time.Now(),
		},
		Title:       title,
		Content:     text,
		ContentHtml: html,
		SortId:      uint(sort),
	}

	err = utils.DB.Create(&article).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": article,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": article,
		}
	}
	c.ServeJSON() // 返回json
}

// GetAll 获取文章列表
func (c *ArticleController) GetAll() {
	var article []models.Article
	err := utils.DB.Order("created_at desc").Find(&article).Error
	if err != nil {
		c.Ctx.WriteString("文章获取失败")
	}
	c.Data["Article"] = article
	c.TplName = "admin/article/article.html"
}

// EditWrite 渲染修改文章页面
func (c *ArticleController) EditWrite() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id错误")
		return
	}
	var article models.Article
	utils.DB.Where("id = ?", id).Find(&article)
	var sort []models.Sort
	err = utils.DB.Find(&sort).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
		return
	}
	c.Data["Sort"] = sort
	c.Data["Result"] = article
	c.TplName = "admin/article/edit_article.html"
}

// EditWritePost 修改文章
func (c *ArticleController) EditWritePost() {
	title := c.GetString("title")
	html := c.GetString("html")
	text := c.GetString("text")
	sort := c.GetString("sort")
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id错误")
		return
	}
	article := models.Article{
		Model: models.Model{
			Id: uint(id),
		},
	}
	err = utils.DB.Model(&article).Updates(map[string]interface{}{"title": title, "content_html": html, "content": text, "sort_id": sort}).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": article,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": article,
		}
	}
	c.ServeJSON() // 返回json
}

// DelWrite 删除文章
func (c *ArticleController) DelWrite() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id错误")
		return
	}
	article := models.Article{
		Model: models.Model{
			Id: uint(id),
		},
	}
	err = utils.DB.Delete(&article).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": article,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": article,
		}
	}
	c.ServeJSON() // 返回json
}
