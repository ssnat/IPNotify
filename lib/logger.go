package lib

import (
	"fmt"
	"log"
	"time"
)

type ILogger struct {
}

func (own *ILogger) Println(loggerType string, content interface{}) {
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", date, loggerType, content)
}

func (own *ILogger) Error(err interface{}) {
	own.Println("Error", err)
}

func (own *ILogger) Info(content interface{}) {
	own.Println("Info", content)
}

func (own *ILogger) PanicError(err error) {
	own.Error(err)
	log.Fatal(err)
}

var Logger = ILogger{}
