//Package routes
/*
@Title: api.go
@Description
@Author: kkw 2023/1/3 15:48
*/
package routes

import (
	"github.com/gin-gonic/gin"
	"go.kkw.top/gamesTimeLine/internal/app/issue"
)

func RegisterAPIRoutes(engine *gin.Engine) {
	var v1 *gin.RouterGroup

	v1 = engine.Group("v1")
	{
		issueGroup := v1.Group("issue")
		{
			issueGroup.POST("/web_hook", issue.WebHook)
		}

	}
}
