package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
	log "github.com/sirupsen/logrus"
)

func init() {
	// // 209
	// path := "./ChargeLog.log"
	// 206
	path := "./Log.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Duration(60 * 60 * 24)*time.Second),
	)
	log.SetOutput(writer)
}