// logger.go
// author:昌维 [github.com/cw1997]
// date:2017-05-11 16:29:22
package logger

import (
	//	"fmt"
	"log"
	"os"
	"time"

	"config"
	"util"
)

var (
	logPath string
	logFile *os.File
	logger  *log.Logger
)

func InitLog() error {
	logPath = getLogPath()
	//	logFile, err := os.OpenFile(logPath+"/"+getLogFileName(), os.O_RDWR|os.O_CREATE, 0)
	logFile, err := os.Create("test.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	logger = log.New(logFile, "", log.LstdFlags|log.Llongfile)
	return err
}

func getLogFileName() string {
	timestamp := util.GetTimestamp()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006_01_02") + ".log"
}

func getLogPath() string {
	if logFilePath := config.Get("log.path"); logFilePath == "" {
		return util.GetCurrentDirectory() + "/log"
	} else {
		return logFilePath
	}
}

func Fatal(msg interface{}) {
	logger.Fatalln(msg)
}

func Print(msg interface{}) {
	logger.Println(msg)
}

func Fatalf(msg string, err error) {
	logger.Fatalf(msg, err)
}
