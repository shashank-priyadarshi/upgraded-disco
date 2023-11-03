package logger

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

type Logger interface {
	Init()
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	With(key, value string) Logger
}

type LoggingService struct {
	Level    string
	Filename string
	file     *os.File
	mux      *sync.RWMutex
}

func NewLogger(level, filename string) Logger {
	l := &LoggingService{
		Level:    level,
		Filename: filename,
		mux:      &sync.RWMutex{},
	}
	l.Init()
	return l
}

func (l *LoggingService) Init() {
	switch l.Level {
	case "debug":
		zap.LevelFlag(l.Filename, zap.DebugLevel, "")
	case "error":
		zap.LevelFlag(l.Filename, zap.ErrorLevel, "")
	case "info":
		zap.LevelFlag(l.Filename, zap.InfoLevel, "")
	default:
		zap.LevelFlag(l.Filename, zap.ErrorLevel, "")
	}
	l.file = os.Stdout
}

func (l *LoggingService) Debugf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	zap.L().Debug(fmt.Sprintf(format, v...))
}

func (l *LoggingService) Infof(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	zap.L().Info(fmt.Sprintf(format, v...))
}

func (l *LoggingService) Errorf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	zap.L().Error(fmt.Sprintf(format, v...))
}

func (l *LoggingService) With(key, value string) Logger {
	keys := map[string]bool{}
	if _, ok := keys[key]; !ok {
		panic("wrong key passed in logger.With()")
	}
	// TODO: enable submodule logging
	return l
}
