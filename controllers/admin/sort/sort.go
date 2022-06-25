package SortController

import (
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type SortController struct {
	beego.Controller
}

// Get 获取分类
func (c *SortController) Get() {
	var sort []models.Sort
	err := utils.DB.Find(&sort).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
		return
	}
	c.Data["Sort"] = sort
	c.TplName = "admin/sort/sort.html"
}

// AddSort 渲染添加分类
func (c *SortController) AddSort() {
	c.TplName = "admin/sort/add_sort.html"
}

// AddSortPost 添加分类的post
func (c *SortController) AddSortPost() {
	name := c.GetString("name")
	if name == "" {
		c.Ctx.WriteString("分类名称不能为空")
		return
	}
	sort := models.Sort{
		Model: models.Model{
			CreatedAt: time.Now(),
		},
		Title: name,
	}
	err := utils.DB.Create(&sort).Error
	if err != nil {
		c.Ctx.WriteString("添加失败")
	} else {
		c.Redirect("/admin/sort", 301)
	}
}

// EditSort 渲染 编辑分类
func (c *SortController) EditSort() {
	id := c.GetString("id")
	var sort models.Sort
	utils.DB.Where("id = ?", id).Find(&sort)
	c.Data["Result"] = sort
	c.TplName = "admin/sort/edit_sort.html"
}

// EditSortPost 编辑分类的post
func (c *SortController) EditSortPost() {
	name := c.GetString("name")
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id错误")
		return
	}
	if name == "" {
		c.Ctx.WriteString("分类名称不能为空")
		return
	}
	sort := models.Sort{
		Model: models.Model{
			Id: uint(id),
		},
	}
	err = utils.DB.Model(&sort).Updates(map[string]interface{}{"title": name}).Error
	if err != nil {
		c.Ctx.WriteString("编辑失败")
	} else {
		c.Redirect("/admin/sort", 301)
	}
}

// DelSort 删除分类
func (c *SortController) DelSort() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id错误")
		return
	}
	sort := models.Sort{
		Model: models.Model{
			Id: uint(id),
		},
	}
	err = utils.DB.Delete(&sort).Error
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": sort,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": sort,
		}
	}
	c.ServeJSON()
}
