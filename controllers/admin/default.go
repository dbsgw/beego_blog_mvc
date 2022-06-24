package admin

import (
	"beego_blog_mvc/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	fmt.Println(utils.GetDay(), utils.GetUnix(), utils.GetUnixNano())
	utils.DB.Raw("show databases")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Arr"] = [5]int{1, 2, 3, 4, 5}
	c.TplName = "admin/index.html"
}

// 登录发token的
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
	result := []map[string]interface{}{}
	utils.DB.Raw("SELECT * FROM users WHERE username=? AND password =?", username, password).Find(&result)
	if len(result) > 0 {
		// 有内容
		Secrect, _ := beego.AppConfig.String("Secrect")
		token, _ := utils.CreateToken(username, Secrect)

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
	c.Ctx.SetCookie("token","",-1)
	c.Redirect("/",301)
}

