package routers

import (
	"beego_blog_mvc/controllers/admin"
	"beego_blog_mvc/controllers/web"
	"beego_blog_mvc/controllers/web/category"
	"beego_blog_mvc/controllers/web/sort"
	"beego_blog_mvc/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 博客前台
	beego.Router("/", &web.MainController{})
	// 分类详情
	beego.Router("/sort/:id", &sort.SortController{})
	// 文章详情
	beego.Router("/category/:id", &category.CategoryController{}, "get:Category")

	// 博客后台api
	var nsAdmin = beego.NewNamespace("/admin",
		//中间件:匹配路由前会执,可以用于权限验证
		//beego.NSBefore(middleware.AdminAuth),

		beego.NSBefore(middleware.AdminAuth),
		// 登录login
		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/logout", &admin.LoginController{}, "get:Logout"),

		beego.NSRouter("/home", &admin.HomeController{}),
		beego.NSRouter("/write", &admin.WriteController{}),
		beego.NSRouter("/addWrite", &admin.WriteController{}, "post:AddWrite"),
		beego.NSRouter("/admin", &admin.AdminController{}),
		beego.NSRouter("/EditWrite", &admin.AdminController{}, "get:EditWrite"),
		beego.NSRouter("/EditWrite", &admin.AdminController{}, "post:EditWritePost"),
		beego.NSRouter("/delWrite", &admin.AdminController{}, "post:DelWrite"),

		// 分类
		beego.NSRouter("/sort", &admin.SortController{}),
		beego.NSRouter("/addSort", &admin.SortController{}, "get:AddSort"),
		beego.NSRouter("/addSort", &admin.SortController{}, "post:AddSortPost"),
		beego.NSRouter("/editSort", &admin.SortController{}, "get:EditSort"),
		beego.NSRouter("/editSort", &admin.SortController{}, "post:EditSortPost"),
		beego.NSRouter("/delSort", &admin.SortController{}, "post:DelSort"),
	)
	beego.AddNamespace(nsAdmin)
}
