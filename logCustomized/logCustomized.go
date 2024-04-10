package logCustomized

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

/*
	func getEncoder() zapcore.Encoder {
		return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	}
*/

func GetLogWriterSingleFile(part, fileName string) zapcore.WriteSyncer {

	file, _ := os.Create(part + fileName)
	return zapcore.AddSync(file)
}

func GetLogWriterSplitFile(part, fileName string) zapcore.WriteSyncer {

	l := &lumberjack.Logger{
		Filename:   part + fileName, //Filename 是要写入日志的文件。
		MaxSize:    1,               //MaxSize 是日志文件在轮换之前的最大大小（以兆字节为单位）。它默认为 100 兆字节
		MaxBackups: 1,               //MaxBackups 是要保留的最大旧日志文件数。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
		MaxAge:     30,              //MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
		Compress:   true,            //压缩
		LocalTime:  true,            //LocalTime 确定用于格式化备份文件中的时间戳的时间是否是计算机的本地时间。默认是使用 UTC 时间。
	}
	return zapcore.AddSync(l)

}

func GetLogConsole() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

func InitLogger(encoder zapcore.Encoder, writeSyncer zapcore.WriteSyncer) *zap.SugaredLogger {

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	//logger := zap.New(core)
	return logger.Sugar()
}
