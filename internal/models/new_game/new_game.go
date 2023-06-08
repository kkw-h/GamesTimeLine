//Package new_game
/*
@Title: new_game.go
@Description
@Author: kkw 2023/6/8 17:21
*/
package new_game

import (
	"go.kkw.top/gamesTimeLine/internal/models"
	"time"
)

var (
	NewGameId        int64 = 5595314348
	NewGameIdString        = "new_game"
	NewVersionId     int64 = 5595312277
	NewVersionString       = "new_version"
)

type NewGame struct {
	ID          int64  `gorm:"uniqueIndex"`
	Name        string `gorm:"type:varchar(300)"`
	ReleaseDate *time.Time
	GenUrl      string
	Desc        string
	IssueUrl    string
	models.ModelDate
}
