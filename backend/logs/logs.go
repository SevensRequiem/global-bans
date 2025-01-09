package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var logfile *os.File

type LogEntry struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Timestamp int64  `json:"timestamp"`
}

type HTTPLogEntry struct {
	Level     string `json:"level"`
	Method    string `json:"method"`
	URI       string `json:"uri"`
	Status    int    `json:"status"`
	UserAgent string `json:"user_agent"`
	IP        string `json:"ip"`
	Time      int64  `json:"time"`
}

func init() {
	var err error
	logfile, err = os.OpenFile("anet.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open log file")
	}
}

func LogInfo(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "INFO",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func LogDebug(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "DEBUG",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func LogError(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "ERROR",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func LogFatal(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "FATAL",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func LogCritical(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "CRITICAL",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func LogHTTP(msg string, line int, file string) {
	entry := LogEntry{
		Level:     "HTTP",
		Message:   msg,
		File:      file,
		Line:      line,
		Timestamp: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(entry)
	go jsonWrite(jsonData)
}

func HttpEchoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		entry := HTTPLogEntry{
			Level:     "HTTP",
			Method:    c.Request().Method,
			URI:       c.Request().URL.Path,
			Status:    c.Response().Status,
			UserAgent: c.Request().UserAgent(),
			IP:        c.RealIP(),
			Time:      time.Now().Unix(),
		}
		jsonData, _ := json.Marshal(entry)
		go jsonWrite(jsonData)
		return err
	}
}

func jsonWrite(data interface{}) {
	logfile.Write(data.([]byte))
	logfile.Write([]byte("\n"))
}
