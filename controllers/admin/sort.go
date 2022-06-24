package admin

import (
	"beego_blog_mvc/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type SortController struct {
	beego.Controller
}

// 获取分类
func (c *SortController) Get() {
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("select * from sort").Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("分类获取失败")
	}
	c.Data["Arr"] = results
	c.TplName = "admin/sort.html"
}

// 渲染 添加分类
func (c *SortController) AddSort() {
	c.TplName = "admin/add_sort.html"
}

// 添加分类的post
func (c *SortController) AddSortPost() {
	name := c.GetString("name")
	if name == "" {
		c.Ctx.WriteString("分类名称不能为空")
		return
	}
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("INSERT INTO sort(title, add_time) VALUES (?, ?)", name, utils.GetUnix()).Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("添加失败")
	} else {
		fmt.Println(name, "添加成功", utils.GetUnix())
		c.Redirect("/admin/sort", 301)
	}
	fmt.Println(err, results, "===")
}

// 渲染 编辑分类
func (c *SortController) EditSort() {
	id := c.GetString("id")
	result := map[string]interface{}{}
	utils.DB.Raw("select * from sort where id = ?", id).Find(&result)
	fmt.Println(result, "id---")

	c.Data["Result"] = result
	c.TplName = "admin/edit_sort.html"
}

// 编辑分类的post
func (c *SortController) EditSortPost() {
	name := c.GetString("name")
	id := c.GetString("id")
	if name == "" {
		c.Ctx.WriteString("分类名称不能为空")
		return
	}
	results := []map[string]interface{}{}
	err := utils.DB.Debug().Raw("update sort set title = ? where id = ?;", name, id).Find(&results).Error
	if err != nil {
		c.Ctx.WriteString("编辑失败")
	} else {
		fmt.Println(name, "编辑成功", utils.GetUnix())
		c.Redirect("/admin/sort", 301)
	}
	fmt.Println(err, results, "===")
}

// 删除分类
func (c *SortController) DelSort() {
	id := c.GetString("id")
	result := map[string]interface{}{}
	err := utils.DB.Raw("delete from sort where id = ?", id).Find(&result).Error
	fmt.Println(result, "id---")

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 20001,
			"msg":  "创建失败",
			"data": result,
		}
	}else {
		c.Data["json"] = map[string]interface{}{
			"code": 20000,
			"msg":  "创建成功",
			"data": result,
		}
	}
	c.ServeJSON() // 返回json
}