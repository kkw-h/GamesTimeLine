//Package setupRouter
/*
@Title: setupRouter.go
@Description
@Author: kkw 2023/1/3 16:57
*/
package setupRouter

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go.kkw.top/gamesTimeLine/internal/config"
	"go.kkw.top/gamesTimeLine/internal/migrates"
	"go.kkw.top/gamesTimeLine/internal/routes"
	"go.kkw.top/gamesTimeLine/pkg"
	"go.kkw.top/gamesTimeLine/pkg/database"
	"go.kkw.top/gamesTimeLine/pkg/redis"
)

func SetupRouter() *gin.Engine {
	//初始化日志
	pkg.New()

	r := gin.Default()
	//加载配置
	var configPath string
	flag.StringVar(&configPath, "config", "./", "配置文件路径")
	config.LoadConfig(configPath)

	if config.Config.Get("databases") != nil {
		//创建数据库链接
		database.Client()
		//初始化数据库
		migrates.Migrate()
	}
	//判断是否加载Redis
	if config.Config.Get("redis") != nil {
		redis.NewClient()
		defer redis.KwClient.Client.Close()
	}

	//加载路由
	routes.RegisterAPIRoutes(r)

	return r
}
