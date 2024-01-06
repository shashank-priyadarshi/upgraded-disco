package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
)

type Logger interface {
	init()
	calculateRelativePath(caller zapcore.EntryCaller) string
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	With(key, value string) Logger
	WithSubmodule(name string) Logger
}

type LoggingService struct {
	Level      string
	context    context.Context
	basePath   string
	file       *os.File
	mux        *sync.RWMutex
	zapLogger  *zap.Logger
	submodules map[string]*LoggingService
	//defaultLog Logger
}

func NewLogger(level, basePath string) Logger {

	l := &LoggingService{
		Level:      level,
		basePath:   basePath,
		mux:        &sync.RWMutex{},
		submodules: make(map[string]*LoggingService),
	}
	l.init()

	//l.defaultLog = l.With("main", "default")

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

	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeName = zapcore.FullNameEncoder
	config.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// Calculate relative path based on the current working directory
		relPath, err := filepath.Rel(l.basePath, caller.TrimmedPath())
		if err != nil {
			// Handle the error if necessary
			relPath = caller.TrimmedPath()
		}
		enc.AppendString(relPath)
	}

	logger, err := config.Build(
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				NewColorEncoder(config.EncoderConfig),
				os.Stdout,
				config.Level,
			)
		}),
	)

	if err != nil {
		panic(err)
	}

	logger = logger.WithOptions(zap.AddCaller())
	logger = logger.WithOptions(zap.AddCallerSkip(1))

	l.zapLogger = logger

}

func (l *LoggingService) calculateRelativePath(caller zapcore.EntryCaller) string {
	relPath, err := filepath.Rel(l.basePath, caller.TrimmedPath())
	if err != nil {
		// Handle the error if necessary
		relPath = caller.TrimmedPath()
	}
	return relPath
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

//func (l *LoggingService) With(key, value string) Logger {
//	newZapLogger := l.zapLogger.With(zap.String(key, value))
//
//	return &LoggingService{
//		Level:     l.Level,
//		context:   l.context,
//		file:      l.file,
//		mux:       l.mux,
//		zapLogger: newZapLogger,
//	}
//}

func (l *LoggingService) With(key, value string) Logger {

	// Create a new LoggingService for the current key-value pair
	newLogger := &LoggingService{
		Level:      l.Level,
		mux:        l.mux,
		zapLogger:  l.zapLogger.With(zap.String(key, value)),
		submodules: make(map[string]*LoggingService),
		//defaultLog: l.defaultLog,
	}

	// Add the new logger to submodules
	//l.submodules[value] = newLogger

	return newLogger
}

func (l *LoggingService) WithSubmodule(name string) Logger {
	// Check if the submodule already exists
	if submodule, ok := l.submodules[name]; ok {
		return submodule
	}

	// Create a new submodule logger
	newLogger := &LoggingService{
		Level:      l.Level,
		mux:        l.mux,
		zapLogger:  l.zapLogger.With(zap.String("submodule", name)),
		submodules: make(map[string]*LoggingService),
		//defaultLog: l.defaultLog,
	}

	// Add the new submodule logger to submodules
	l.submodules[name] = newLogger

	return newLogger
}
