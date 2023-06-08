//Package issue
/*
@Title: issue.go
@Description
@Author: kkw 2023/6/8 16:29
*/
package issue

import (
	"github.com/gin-gonic/gin"
	"go.kkw.top/gamesTimeLine/internal/models/new_game"
	"go.kkw.top/gamesTimeLine/pkg"
	"net/http"
)

func WebHook(ctx *gin.Context) {
	var req WebHookReq
	// 验证 GitHub Webhook 请求的签名
	signature := ctx.GetHeader("X-Hub-Signature")
	pkg.KwLogger.Info("signature", signature)
	// 进行验证逻辑...

	// 解析请求的 JSON 数据
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// 处理 GitHub Webhook 请求的业务逻辑
	pkg.KwLogger.Info("payload", req)
	//验证标签
	labelStr := _validateLabel(req.Issue.Labels)
	if labelStr == "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
	}

	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
}

func _validateLabel(labels []LabelReq) string {
	if len(labels) > 0 {
		for _, label := range labels {
			if label.ID == new_game.NewGameId {
				//新游发布
				return new_game.NewGameIdString
			}
			if label.ID == new_game.NewVersionId {
				//版本更新
				return new_game.NewVersionString
			}
		}
	}
	pkg.KwLogger.Debug("未解析标签", labels)
	return ""
}
