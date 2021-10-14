package errorlog

import (
	"fmt"
	"os"
	"runtime/debug"
	"teammap/pkg/myerr"
	"teammap/pkg/setting"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Setup() {
	Logger = logrus.New()

	errorLogW, _ := os.OpenFile("./log/error/errors.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	infoLogW, _ := os.OpenFile("./log/error/info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)

	formatter := new(ErrorFormatter)
	Logger.SetFormatter(formatter)
	Logger.Out = os.Stdout

	var writerMap lfshook.WriterMap
	switch setting.AppSetting.RunMode {
	case "release":
		// 线上环境
		Logger.Level = logrus.ErrorLevel
		writerMap = lfshook.WriterMap{
			logrus.PanicLevel: errorLogW,
			logrus.FatalLevel: errorLogW,
			logrus.ErrorLevel: errorLogW,
		}
	case "test":
		// 测试环境info级别输出到文件，debug日志到std
		Logger.Level = logrus.DebugLevel
		writerMap = lfshook.WriterMap{
			logrus.PanicLevel: errorLogW,
			logrus.FatalLevel: errorLogW,
			logrus.ErrorLevel: errorLogW,
			logrus.WarnLevel:  infoLogW,
			logrus.InfoLevel:  infoLogW,
		}
		break
	case "debug":
		// debug环境所有日志不处理，输出到std
		Logger.Level = logrus.DebugLevel
		break
	}

	lfHook := lfshook.NewHook(writerMap, formatter)
	Logger.Hooks.Add(lfHook)
}

type ErrorFormatter struct{}

func (f *ErrorFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s/%s]:%s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		entry.Level.String(),
		entry.Message)
	return []byte(msg), nil
}

func LogMyError(err *myerr.MyErr) {
	if err == nil {
		return
	}
	switch err.Level {
	case myerr.ErrorLevel:
		Logger.Error(err.Detail())
		break
	case myerr.WarnLevel:
		Logger.Warn(err.Detail())
		break
	case myerr.InfoLevel:
		Logger.Info(err.Error())
		break
	case myerr.DebugLevel:
		Logger.Debug(err.Error())
		break
	}
}

func Debug(msg ...interface{}) {
	Logger.Debug(msg...)
}

func Info(msg ...interface{}) {
	Logger.Info(msg...)
}

func Error(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}

func Fatal(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}

func Panic(msg string) {
	msg += string(debug.Stack())
	Logger.Error(msg)
}
