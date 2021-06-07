package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

//初始化日志
func InitLog(level string){
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	if level =="debug"{
		atomicLevel.SetLevel(zap.DebugLevel)
	}else if level=="info"{
		atomicLevel.SetLevel(zap.InfoLevel)
	}else if level=="error"{
		atomicLevel.SetLevel(zap.ErrorLevel)
	}else{
		atomicLevel.SetLevel(zap.DebugLevel)
	}
	atomicLevel.SetLevel(zap.DebugLevel)
	core := zapcore.NewCore(encoder, writeSyncer,atomicLevel)
	//zap.AddCaller() 开启开发模式，堆栈跟踪
	zapLogger := zap.New(core, zap.AddCaller(),zap.Development())
	Logger = zapLogger.Sugar()
	defer Logger.Sync()
}

func getLogWriter() zapcore.WriteSyncer{
	/**
	Filename: 日志文件的位置
	MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups：保留旧文件的最大个数
	MaxAges：保留旧文件的最大天数
	Compress：是否压缩/归档旧文件
	*/
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/ueba-api.log",
		MaxSize:    128, // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,   // 日志文件最多保存多少个备份
		MaxAge:     30,   // 文件最多保存多少天
		Compress:   false,// 是否压缩
	}
	// 打印到控制台和文件
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(lumberJackLogger))
	return writeSyncer
}

//日志编码器配置
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		}, //zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

type CustomLog struct {
	File    string
	Line    int
	Message string
}