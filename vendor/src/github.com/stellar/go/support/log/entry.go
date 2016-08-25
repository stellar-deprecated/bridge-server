package log

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/go/support/errors"
)

func (e *Entry) SetLevel(level logrus.Level) {
	e.Logger.Level = level
}

// WithField creates a child logger annotated with the provided key value pair.
// A subsequent call to one of the logging methods (Debug(), Error(), etc.) to
// the return value from this function will cause the emitted log line to
// include the provided value.
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{*e.Entry.WithField(key, value)}
}

// WithFields creates a child logger annotated with the provided key value
// pairs.
func (e *Entry) WithFields(fields F) *Entry {
	return &Entry{*e.Entry.WithFields(logrus.Fields(fields))}
}

// WithStack annotates this error with a stack trace from `stackProvider`, if
// available.  normally `stackProvider` would be an error that implements
// `errors.StackTracer`.
func (e *Entry) WithStack(stackProvider interface{}) *Entry {
	stack := "unknown"

	if stackProvider, ok := stackProvider.(errors.StackTracer); ok {
		stack = fmt.Sprint(stackProvider.StackTrace())
	}

	return e.WithField("stack", stack)
}

// Debugf logs a message at the debug severity.
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.Entry.Debugf(format, args...)
}

// Debug logs a message at the debug severity.
func (e *Entry) Debug(args ...interface{}) {
	e.Entry.Debug(args...)
}

// Infof logs a message at the Info severity.
func (e *Entry) Infof(format string, args ...interface{}) {
	e.Entry.Infof(format, args...)
}

// Info logs a message at the Info severity.
func (e *Entry) Info(args ...interface{}) {
	e.Entry.Info(args...)
}

// Warnf logs a message at the Warn severity.
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.Entry.Warnf(format, args...)
}

// Warn logs a message at the Warn severity.
func (e *Entry) Warn(args ...interface{}) {
	e.Entry.Warn(args...)
}

// Errorf logs a message at the Error severity.
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.Entry.Errorf(format, args...)
}

// Error logs a message at the Error severity.
func (e *Entry) Error(args ...interface{}) {
	e.Entry.Error(args...)
}

// Panicf logs a message at the Panic severity.
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.Entry.Panicf(format, args...)
}

// Panic logs a message at the Panic severity.
func (e *Entry) Panic(args ...interface{}) {
	e.Entry.Panic(args...)
}
