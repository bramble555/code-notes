package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

/*
logger:
  level: info
  prefix: "[gvb]"
  director: log
  show_line: true
  log_in_console: true
*/

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

// 全局变量
var logger = &Logger{"debug", "pre", "log", true, true}

const (
	red    = 31
	green  = 32
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// 自定义格式实现这个Formatter接口
func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.TraceLevel, logrus.DebugLevel:
		levelColor = blue
	case logrus.InfoLevel, logrus.WarnLevel:
		levelColor = green
	default:
		levelColor = red
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := logger
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// entry.HasCaller() 来决定是否添加调用者信息
	if entry.HasCaller() {
		funVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m  %s %s %s\n",
			log.Prefix,
			timestamp,
			levelColor,
			entry.Level,
			fileVal,
			funVal,
			entry.Message,
		)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m  %s\n",
			log.Prefix,
			timestamp,
			levelColor,
			entry.Level,
			entry.Message,
		)
	}
	return b.Bytes(), nil
}
func createLogFile() *os.File {
	// 输出到哪
	logFilePath := filepath.Join("log", "app.log")

	// 创建文件（如果不存在）
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 这里只是一个示例，实际使用中可能需要更复杂的错误处理
		panic(err)
	}
	return file
}
func InitLogger() *logrus.Logger {
	mLog := logrus.New()

	// 设置输出为控制台和文件
	file := createLogFile()
	mLog.SetOutput(io.MultiWriter(os.Stdout, file))
	// 是否开启行号
	mLog.SetReportCaller(logger.ShowLine)
	// 格式
	mLog.SetFormatter(&LogFormatter{})
	// 解析level
	level, err := logrus.ParseLevel(logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	// 设置level
	mLog.SetLevel(level)
	return mLog
}
func main() {
	log := InitLogger()
	log.Debugln("debug")
	log.Infoln("info")
	log.Fatalln("fatal")
}
