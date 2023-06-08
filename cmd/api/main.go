//Package api
/*
@Title: main.go
@Description
@Author: kkw 2022/12/30 16:52
*/
package main

import (
	"go.kkw.top/gamesTimeLine/cmd/api/setupRouter"
)

func main() {
	r := setupRouter.SetupRouter()
	//启动端口
	err := r.Run()
	if err != nil {
		panic("启动失败")
	} // 监听并在 0.0.0.0:8080 上启动服务
}
