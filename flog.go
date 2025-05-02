package flog

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log           *zap.Logger
	DefaultLogger *zap.Logger
	level         zap.AtomicLevel
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type (
	Level = zapcore.Level
	Field = zap.Field
)

var DebugLevel = zapcore.DebugLevel
var InfoLevel = zapcore.InfoLevel
var WarnLevel = zapcore.WarnLevel
var ErrorLevel = zapcore.ErrorLevel
var FatalLevel = zapcore.FatalLevel

func init() {
	var err error
	c := zap.NewProductionConfig()
	level = zap.NewAtomicLevelAt(DebugLevel)
	c.Encoding = "console"
	c.Level = level
	log, err = c.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}

	cc := zap.NewProductionConfig()
	cc.Encoding = "console"
	cc.Level = level
	DefaultLogger, err = cc.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
}

func SetLevel(lvl Level) {
	level.SetLevel(lvl)
}

func DebugWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

func InfoWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Info(msg, fields...)
}

func WarnWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

func ErrorWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Error(msg, fields...)
}

func FatalWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Fatal(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	log.Fatal(msg, fields...)
}
