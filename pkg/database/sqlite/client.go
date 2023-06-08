// Package sqlite
/*
@Title: db.go
@Description
@Author: kkw 2022/12/30 17:41
*/
package sqlite

import (
	"fmt"
	"go.kkw.top/gamesTimeLine/internal/config"
	"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	// "github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(
		fmt.Sprintf("%s.db", config.Get("databases.sqlite.dbName"))),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true},
	)
}
