package molog

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type Level = logrus.Level

const (
	// panicLevel, fatalLevel are not public because they're unused.
	// They exist just to match the level numbers with logrus.
	panicLevel Level = iota
	fatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type Fields = logrus.Fields

type Formatter = logrus.Formatter
type JSONFormatter = logrus.JSONFormatter
type TextFormatter = logrus.TextFormatter

type Entry struct {
	logger   *Logger
	internal *logrus.Entry
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	limiter := entry.logger.limiter
	if limiter == nil || limiter.Allow(entry) {
		entry.internal.Log(level, args...)
	}
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.Log(TraceLevel, args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.Log(DebugLevel, args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.Log(ErrorLevel, args...)
}

func (entry *Entry) GetFields() Fields {
	return entry.internal.Data
}

type Logger struct {
	internal *logrus.Logger
	limiter  Limiter
}

var (
	defaultLogger   = new(logrus.StandardLogger())
	RequiredLimiter Limiter
)

func New() *Logger {
	return new(logrus.New())
}

func new(internal *logrus.Logger) *Logger {
	return &Logger{
		internal: internal,
	}
}

func applyRequiredLimiter(limiter Limiter) Limiter {
	if RequiredLimiter != nil {
		return AndLimiters(RequiredLimiter, limiter)
	}
	return limiter
}

func (l *Logger) SetFormatter(formatter Formatter) {
	l.internal.SetFormatter(formatter)
}

func (l *Logger) SetOutput(output io.Writer) {
	l.internal.SetOutput(output)
}

func (l *Logger) Limited(limiter Limiter) *Logger {
	return &Logger{
		internal: l.internal,
		limiter:  applyRequiredLimiter(limiter),
	}
}

func (l *Logger) newEntry(internal *logrus.Entry) *Entry {
	return &Entry{
		logger:   l,
		internal: internal,
	}
}

func (l *Logger) WithField(key string, value any) *Entry {
	return l.newEntry(l.internal.WithField(key, value))
}

func (l *Logger) WithFields(fields Fields) *Entry {
	return l.newEntry(l.internal.WithFields(fields))
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	entry := l.newEntry(logrus.NewEntry(l.internal))
	entry.Log(level, fmt.Sprintf(format, args...))
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Logf(TraceLevel, format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logf(DebugLevel, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logf(WarnLevel, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}

func SetFormatter(formatter Formatter) {
	defaultLogger.internal.SetFormatter(formatter)
}

func SetOutput(output io.Writer) {
	defaultLogger.internal.SetOutput(output)
}

func Limited(limiter Limiter) *Logger {
	return defaultLogger.Limited(limiter)
}

func WithField(key string, value any) *Entry {
	return defaultLogger.WithField(key, value)
}

func WithFields(fields Fields) *Entry {
	return defaultLogger.WithFields(fields)
}

func Logf(level Level, format string, args ...interface{}) {
	defaultLogger.Logf(level, format, args...)
}

func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}
