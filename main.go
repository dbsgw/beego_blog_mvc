package main

import (
	_ "beego_blog_mvc/routers"
	"beego_blog_mvc/utils"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.AddFuncMap("index",index)
	beego.AddFuncMap("UnixToDate",UnixToDate)
	beego.Run()
}
// 时间戳转日期
func UnixToDate(unix int64) string  {
	return utils.UnixDate(unix)
}
// 序号加1  $index + 1
func index(in int) (out int) {
	out = in + 1
	return out
}
