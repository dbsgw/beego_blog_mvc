package middleware

import (
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"net/url"
)

func AdminAuth(ctx *context2.Context) {
	pathname := ctx.Request.URL.String()
	urlPath, _ := url.Parse(pathname) //urlPath.Path  /role/edit
	if string(urlPath.Path) != "/admin/login" {
		cookie := ctx.Input.Cookie("token")
		Secret, _ := beego.AppConfig.String("Secret")
		_, err := ParseToken(cookie, Secret)
		if err != nil {
			ctx.SetCookie("token", "", -1)
			ctx.Redirect(302, "/admin/login")
		}
	}
}

// ParseToken 解析token
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
