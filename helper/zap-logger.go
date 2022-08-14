package helper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// var LogObj *zap.Logger
var SugarObj *zap.SugaredLogger
var LogerObj *zap.Logger

type SugaredLogger struct {
	sugarObj *zap.SugaredLogger
}

func init() {

	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   "logs/access.log",
		MaxSize:    1,   // Max megabytes before log is rotated
		MaxBackups: 90,  // Max number of old log files to keep
		MaxAge:     180, // Max number of days to retain log files
		Compress:   true,
	}
	writerSyncer := zapcore.AddSync(lumberjackLogrotate)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	atom := zap.NewAtomicLevel()
	LogerObj = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(writerSyncer),
		atom,
	), zap.AddCaller())

	SugarObj = LogerObj.Sugar()

	defer LogerObj.Sync()
}

func (s *SugaredLogger) Debug(args ...interface{}) {
	s.Debug(args...)
}

func (s *SugaredLogger) Info(args ...interface{}) {
	s.Info(args...)
}

func (s *SugaredLogger) Warn(args ...interface{}) {
	s.Warn(args...)
}

func (s *SugaredLogger) Error(args ...interface{}) {
	s.Error(args...)
}
