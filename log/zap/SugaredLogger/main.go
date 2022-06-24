package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var log *zap.Logger
var logs *zap.SugaredLogger

var db *sqlx.DB

func main() {
	InitLogger()
	defer log.Sync()

}

func InitLogger() {
	// 1. writeSyncer循环文件
	// file, _ := os.Create("./test.log")
	// writeSyncer := zapcore.AddSync(file)
	// lumberJackLogger := &lumberjack.Logger{
	// Filename: "./test.log",
	// LocalTime: true,
	// MaxSize: 1, //M
	// MaxBackups: 5, //个
	// MaxAge: 0, //天
	// Compress: false,
	// }
	// writeSyncer := zapcore.AddSync(lumberJackLogger)

	// 2. console 打印 调试
	writeSyncer := zapcore.AddSync(os.Stdout)

	// 3. encoder JSON格式
	// encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 4. encoder Console格式
	// encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// 5. 自定义格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 6.core
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 7.log
	log = zap.New(core)
	log.Info("ok")
	log.Error("ok")
	logs = log.Sugar()
	logs.Info("logs")
}
