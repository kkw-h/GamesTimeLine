// Package migrates
/*
@Title: migrate.go
@Description
@Author: kkw 2023/1/3 16:08
*/
package migrates

import (
	"go.kkw.top/gamesTimeLine/internal/models/new_game"
	"go.kkw.top/gamesTimeLine/pkg/database"
)

func Migrate() {
	//迁移数据库
	err := database.DB.AutoMigrate(&new_game.NewGame{})
	if err != nil {
		panic("初始化数据库失败")
	}
}
