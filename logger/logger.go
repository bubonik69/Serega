package logger

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var LocalLog *log.Logger

const (
	fileLocalLog = "logger/tmp/local.log"
	sizeLocalLog = 104857600
)

func InitLogger(fileLog string, size int64, isFill bool) *log.Logger {
	// set location of log file
	fileInfo, err := os.Stat(fileLog)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Failed to open log file", fileLog, ":", err)
		}
	} else {
		if fileInfo.Size() >= size {
			os.Remove(fileLog)
			os.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		}
	}
	flag.Parse()

	file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file ", fileLog, " : ", err)
		os.Exit(1)
	}
	multi := io.MultiWriter(file, os.Stdout)
	if isFill {
		return log.New(multi, "", log.LstdFlags|log.Llongfile)
	} else {
		return log.New(multi, "", 0)
	}
}

func InitLocalLogger() {
	LocalLog = InitLogger(fileLocalLog, sizeLocalLog, true)
}