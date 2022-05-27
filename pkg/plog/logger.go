package plog

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

// InitFileLogger
// logPath 需要绝对路径,
// logName 文件名
// accessLogName
func InitFileLogger(logPath, logName string) {
	setFile(logger, logPath, logName)
}

// 不要标准输出
func setOutputNull(l *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Open Src File err", err)
	}
	writer := bufio.NewWriter(src)
	l.SetOutput(writer)
}


func setFile(l *logrus.Logger, logPath, logName string) {
	setOutputNull(l)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		path.Join(logPath, logName)+".%Y%m%d",
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(3*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	l.AddHook(NewFilenameHook())
	l.AddHook(lfHook)
}

func InfoWithFields(arg logrus.Fields) {
	logger.WithFields(arg).Info()
}

func ErrorWithFields(arg logrus.Fields) {
	logger.WithFields(arg).Error()
}

func WarnWithFields(arg logrus.Fields) {
	logger.WithFields(arg).Warn()
}

func WithFields(arg logrus.Fields) *logrus.Entry {
	return logger.WithFields(arg)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(str string, args ...interface{}) {
	logger.Infof(str, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(str string, args ...interface{}) {
	logger.Warnf(str, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(str string, args ...interface{}) {
	logger.Errorf(str, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(str string, args ...interface{}) {
	logger.Fatalf(str, args...)
}

func Printf(str string, args ...interface{}) {
	logger.Infof(str, args...)
}

func Println(args ...interface{}) {
	logger.Info(args...)
}
