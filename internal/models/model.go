//Package models
/*
@Title: model.go
@Description
@Author: kkw 2023/5/25 14:20
*/
package models

import (
	"gorm.io/gorm"
	"time"
)

type ModelDate struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
