//Package mysql
/*
@Title: client.go
@Description
@Author: kkw 2023/5/15 10:31
*/
package mysql

import (
	"go.kkw.top/gamesTimeLine/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Get("databases.mysql.source")), &gorm.Config{})
}
