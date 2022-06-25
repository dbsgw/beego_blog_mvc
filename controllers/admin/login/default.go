package LoginController

import (
	"beego_blog_mvc/models"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "admin/login/index.html"
}

// Post 登录发token的
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Data["json"] = map[string]interface{}{
			"code": false,
			"data": "用户名密码不能为空",
			"msg":  "用户名密码不能为空",
		}
		c.ServeJSON()
		return
	}
	var user []models.User
	utils.DB.Where("username=? AND password =?", username, password).Find(&user)
	if len(user) > 0 {
		// 有内容
		Secret, _ := beego.AppConfig.String("Secret")
		token, _ := utils.CreateToken(username, Secret)
		c.Data["json"] = map[string]interface{}{
			"code":  true,
			"token": token,
			"msg":   "登录成功",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": false,
			"data": "用户名密码错误",
			"msg":  "用户名密码错误",
		}
	}
	c.ServeJSON()
}

func (c *LoginController) Logout() {
	c.Ctx.SetCookie("token", "", -1)
	c.Redirect("/", 301)
}
