package goutil

import (
	"fmt"
	"log"
	"os"
)

var slog = log.New(os.Stdout, "", log.LstdFlags)
var elog = log.New(os.Stderr, "", log.LstdFlags)

func LogInfo(source, format string, v ...interface{}) {
	slog.Printf(fmt.Sprintf("[INFO] %s %s", source, format), v...)
}

func LogDebug(source, format string, v ...interface{}) {
	slog.Printf(fmt.Sprintf("[DEBUG] %s %s", source, format), v...)
}

func LogError(source, format string, v ...interface{}) {
	elog.Printf(fmt.Sprintf("[DEBUG] %s %s", source, format), v...)
}

func LogErr(source string, err error) {
	elog.Println("[ERROR]", source, err)
}

func ErrExit(err error) {
	LogErr("[APP]", err)
	os.Exit(1)
}
