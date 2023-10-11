package creational

import (
	"fmt"
	"sync"
)

type Logger struct {
	Log []string
}

var instance *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) AddLog(log string) {
	l.Log = append(l.Log, log)
}

func (l *Logger) ShowLogs() {
	for _, log := range l.Log {
		fmt.Println(log)
	}
}
