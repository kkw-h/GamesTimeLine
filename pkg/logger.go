// Package pkg
/*
@Title: logger.go
@Description
@Author: kkw 2023/5/15 11:34
*/
package pkg

import (
	log "github.com/sirupsen/logrus"
)

var KwLogger *log.Logger

func New() {
	KwLogger = log.New()
	KwLogger.SetLevel(log.TraceLevel)
}
