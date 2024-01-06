package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"

	"go.uber.org/zap"
)

type Logger interface {
	init()
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	With(key, value string) Logger
}

type LoggingService struct {
	Level     string
	context   context.Context
	Filename  string
	file      *os.File
	mux       *sync.RWMutex
	zapLogger *zap.Logger
}

func NewLogger(level, filename string) Logger {

	l := &LoggingService{
		Level:    level,
		Filename: filename,
		mux:      &sync.RWMutex{},
	}
	l.init()

	return l
}

func (l *LoggingService) init() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	switch l.Level {
	case "DEBUG":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "ERROR":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "INFO":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	//config.OutputPaths = []string{"logs.txt"}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	logger = logger.WithOptions(zap.AddCaller())
	logger = logger.WithOptions(zap.AddCallerSkip(1))

	l.zapLogger = logger

}

func (l *LoggingService) Debugf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.zapLogger.Debug(fmt.Sprintf(format, v...))
}

func (l *LoggingService) Infof(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.zapLogger.Info(fmt.Sprintf(format, v...))
}

func (l *LoggingService) Errorf(format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.zapLogger.Error(fmt.Sprintf(format, v...))
}

func (l *LoggingService) With(key, value string) Logger {
	keys := map[string]bool{}
	if _, ok := keys[key]; !ok {
		panic("wrong key passed in logger.With()")
	}
	// TODO: enable submodule logging
	return l
}
