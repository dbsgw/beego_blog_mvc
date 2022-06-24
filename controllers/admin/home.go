package admin

import (
	"beego_blog_mvc/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	cookie := c.Ctx.Input.Cookie("token")
	fmt.Println(cookie)
	utils.DB.Raw("show databases")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Arr"] = [5]int{1, 2, 3, 4, 5}
	c.TplName = "admin/home.html"
}
