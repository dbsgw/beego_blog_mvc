package routers

import (
	"beego_blog_mvc/controllers/admin"
	ArticleController "beego_blog_mvc/controllers/admin/article"
	LoginController "beego_blog_mvc/controllers/admin/login"
	adminSort "beego_blog_mvc/controllers/admin/sort"
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
	beego.Router("/sort/:id", &SortController.SortController{})
	// 文章详情
	beego.Router("/category/:id", &CategoryController.CategoryController{}, "get:Category")

	// 博客后台api
	var nsAdmin = beego.NewNamespace("/admin",
		//中间件:匹配路由前会执,可以用于权限验证
		//beego.NSBefore(middleware.AdminAuth),

		beego.NSBefore(middleware.AdminAuth),
		// 登录login
		beego.NSRouter("/login", &LoginController.LoginController{}),
		beego.NSRouter("/logout", &LoginController.LoginController{}, "get:Logout"),

		// 文章管理
		beego.NSRouter("/home", &admin.HomeController{}),
		beego.NSRouter("/write", &ArticleController.ArticleController{}),
		beego.NSRouter("/addWrite", &ArticleController.ArticleController{}, "post:AddWrite"),
		beego.NSRouter("/admin", &ArticleController.ArticleController{}, "get:GetAll"),
		beego.NSRouter("/EditWrite", &ArticleController.ArticleController{}, "get:EditWrite"),
		beego.NSRouter("/EditWrite", &ArticleController.ArticleController{}, "post:EditWritePost"),
		beego.NSRouter("/delWrite", &ArticleController.ArticleController{}, "post:DelWrite"),

		// 分类管理
		beego.NSRouter("/sort", &adminSort.SortController{}),
		beego.NSRouter("/addSort", &adminSort.SortController{}, "get:AddSort"),
		beego.NSRouter("/addSort", &adminSort.SortController{}, "post:AddSortPost"),
		beego.NSRouter("/editSort", &adminSort.SortController{}, "get:EditSort"),
		beego.NSRouter("/editSort", &adminSort.SortController{}, "post:EditSortPost"),
		beego.NSRouter("/delSort", &adminSort.SortController{}, "post:DelSort"),
	)
	beego.AddNamespace(nsAdmin)
}
