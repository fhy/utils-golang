package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fhy/utils-golang/config"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	logger "github.com/sirupsen/logrus"
)

type webbFormatter struct {
}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (mf *webbFormatter) Format(entry *logger.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	pathList := strings.Split(entry.Caller.File, "/")
	fileName := ""
	if len(pathList) > 1 {
		fileName = pathList[len(pathList)-2] + "/"
	}
	fileName += pathList[len(pathList)-1]
	b.WriteString(fmt.Sprintf("[%s]-[%s:%d]-[%s] - %s\n",
		entry.Time.Format("2006-01-02 15:04:05.0000"), fileName,
		entry.Caller.Line, level, entry.Message))
	return b.Bytes(), nil
}

func InitLogger(logConfig *config.LogConfig) {
	path := logConfig.Path
	if len(path) >= 0 {
		dir := filepath.Dir(path)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if os.MkdirAll(dir, 0755) != nil {
				fmt.Printf("failed to create dir %s, error: %s", dir, err)
			}
		}
	} else {
		path, _ = os.Getwd()
		path += filepath.Base(os.Args[0]) + ".log"
	}
	rotateCount := logConfig.RotationCount
	if rotateCount == 0 {
		rotateCount = 5
	}
	writter, _ := rotatelogs.New(path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationCount(rotateCount),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logger.SetOutput(writter)
	logger.SetFormatter(&webbFormatter{})
	logger.SetReportCaller(true)
	level, err := logger.ParseLevel(logConfig.Level)
	if err != nil {
		level = logger.InfoLevel
	}
	logger.SetLevel(level)
}
