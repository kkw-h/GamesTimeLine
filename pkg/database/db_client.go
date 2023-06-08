// Package database
/*
@Title: client.go
@Description
@Author: kkw 2023/5/15 10:34
*/
package database

import (
	"fmt"
	"go.kkw.top/gamesTimeLine/internal/config"
	"go.kkw.top/gamesTimeLine/pkg/database/mysql"
	"go.kkw.top/gamesTimeLine/pkg/database/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Client() {
	var err error
	switch config.Get("databases.type") {
	case "sqlite":
		DB, err = sqlite.NewClient()
		break
	case "mysql":
		DB, err = mysql.NewClient()
		break
	default:
		panic(fmt.Sprintf("错误的数据库类型:%s", config.Get("databases.type")))
	}

	if err != nil {
		panic(fmt.Sprintf("数据库链接失败, %s", err.Error()))
	}
	SQLDB, err := DB.DB()
	if err != nil {
		panic("数据库链接失败")
	}
	SQLDB.SetMaxOpenConns(config.Config.GetInt("databases.maxOpenConnections"))
	SQLDB.SetMaxIdleConns(config.Config.GetInt("databases.maxIdleConnections"))
	SQLDB.SetConnMaxIdleTime(time.Duration(config.Config.GetInt("databases.maxLifeSeconds")) * time.Second)
}
