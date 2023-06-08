//Package middleware
/*
@Title: jwt.go
@Description
@Author: kkw 2023/5/24 15:51
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"go.kkw.top/gamesTimeLine/internal/response"
	"go.kkw.top/gamesTimeLine/pkg"
	"go.kkw.top/gamesTimeLine/tools/kw_token"
	"net/http"
	"strings"
)

func AuthValidate(isLogin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// 验证 Authorization 请求头是否存在
		if authHeader == "" {
			if isLogin {
				response.ErrorResult(ctx, &response.Error{
					Code:     http.StatusUnauthorized,
					ErrCode:  1401,
					Message:  "登录失效,请重新登录",
					ShowType: response.ShowTypeErr,
				})
				ctx.Abort()
				return
			}
		}

		// 检查 Authorization 请求头是否以 "Bearer " 开头
		if !strings.HasPrefix(authHeader, "Bearer ") {
			if isLogin {
				response.ErrorResult(ctx, &response.Error{
					Code:     http.StatusUnauthorized,
					ErrCode:  1402,
					Message:  "登录失效,请重新登录",
					ShowType: response.ShowTypeErr,
				})
				ctx.Abort()
				return
			}
		}

		// 提取 token 字符串并去除 "Bearer " 前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			if isLogin {
				response.ErrorResult(ctx, &response.Error{
					Code:     http.StatusUnauthorized,
					ErrCode:  1403,
					Message:  "登录失效,请重新登录",
					ShowType: response.ShowTypeErr,
				})
				ctx.Abort()
				return
			}
		}
		tokenContent, err := kw_token.ValidateToken(tokenString)
		if isLogin {
			if err != nil {
				pkg.KwLogger.Debug("解析Token错误", err)
				response.ErrorResult(ctx, &response.Error{
					Code:     http.StatusUnauthorized,
					ErrCode:  1404,
					Message:  "登录失效,请重新登录",
					ShowType: response.ShowTypeErr,
				})
				ctx.Abort()
				return
			}
			ctx.Set("user_id", tokenContent.Id)
			ctx.Set("user_phone", tokenContent.Phone)
		}

		ctx.Next()
	}
}
