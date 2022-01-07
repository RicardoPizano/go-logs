package logger

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getFormatEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFormatEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case infoLevel:
		return zapcore.InfoLevel
	case warnLevel:
		return zapcore.WarnLevel
	case debugLevel:
		return zapcore.DebugLevel
	case errorLevel:
		return zapcore.ErrorLevel
	case fatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	cores := []zapcore.Core{}

	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(), writer, level)
		cores = append(cores, core)
	}

	if config.EnableFile {
		level := getZapLevel(config.FileLevel)
		logf, _ := rotatelogs.New(config.FileLocation+".%Y-%m-%d-%H.log",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithRotationTime(time.Hour),
		)
		writer := zapcore.AddSync(logf)
		core := zapcore.NewCore(getEncoder(), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debug(moduleName string, functionName string, text string) {
	l.sugaredLogger.Debugf("[%s][%s] %s", moduleName, functionName, text)
}

func (l *zapLogger) Info(moduleName string, functionName string, text string) {
	l.sugaredLogger.Infof("[%s][%s] %s", moduleName, functionName, text)
}

func (l *zapLogger) Warn(moduleName string, functionName string, text string) {
	l.sugaredLogger.Warnf("[%s][%s] %s", moduleName, functionName, text)
}

func (l *zapLogger) Error(moduleName string, functionName string, text string) {
	l.sugaredLogger.Errorf("[%s][%s] %s", moduleName, functionName, text)
}

func (l *zapLogger) Fatal(moduleName string, functionName string, text string) {
	l.sugaredLogger.Fatalf("[%s][%s] %s", moduleName, functionName, text)
}

func (l *zapLogger) Request(method string, statusCode int, uri string, start time.Time) {
	l.sugaredLogger.Infof("REQS [%s][%d] %s\t%.2f ms", method, statusCode, uri, float32(time.Since(start).Nanoseconds())/1000000.0)
}
