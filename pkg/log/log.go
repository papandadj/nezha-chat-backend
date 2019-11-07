package log

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Level describes the log severity level.
type Level uint8

var (
	//workspace defined idea path, to easy location code .
	workspace string
	//rootPackageSlash defined file root package slash number.
	//for example  github.com/papandadj/nezha-chat-backend set rootPackageSlash number 3. log will only print
	//file after root package
	rootPackageSlash int
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// Logger is an interface that describes logging.
type Logger interface {
	SetWorkspace(path string)
	SetRootPackageSlash(number int)

	SetLevel(level Level)
	SetOut(out io.Writer)

	Debug(...interface{})
	Debugln(...interface{})

	Info(...interface{})
	Infoln(...interface{})

	Warn(...interface{})
	Warnln(...interface{})

	Error(...interface{})
	Errorln(...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})

	With(key string, value interface{}) Logger
	WithError(err error) Logger
}

type logger struct {
	entry *logrus.Entry
}

// With attaches a key-value pair to a logger.
func (l logger) With(key string, value interface{}) Logger {
	return logger{l.entry.WithField(key, value)}
}

// WithError attaches an error to a logger.
func (l logger) WithError(err error) Logger {
	return logger{l.entry.WithError(err)}
}

// SetLevel sets the level of a logger.
func (l logger) SetLevel(level Level) {
	l.entry.Logger.Level = logrus.Level(level)
}

// SetOut sets the output destination for a logger.
func (l logger) SetOut(out io.Writer) {
	l.entry.Logger.Out = out
}

// Debug logs a message at level Debug on the standard logger.
func (l logger) Debug(args ...interface{}) {
	l.sourced().Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (l logger) Debugln(args ...interface{}) {
	l.sourced().Debugln(args...)
}

// Info logs a message at level Info on the standard logger.
func (l logger) Info(args ...interface{}) {
	l.sourced().Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (l logger) Infoln(args ...interface{}) {
	l.sourced().Infoln(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logger) Warn(args ...interface{}) {
	l.sourced().Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (l logger) Warnln(args ...interface{}) {
	l.sourced().Warnln(args...)
}

// Error logs a message at level Error on the standard logger.
func (l logger) Error(args ...interface{}) {
	l.sourced().Error(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (l logger) Errorln(args ...interface{}) {
	l.sourced().Errorln(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l logger) Fatal(args ...interface{}) {
	l.sourced().Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func (l logger) Fatalln(args ...interface{}) {
	l.sourced().Fatalln(args...)
}

// sourced adds a source field to the logger that contains
// the file name and line where the logging happened.
func (l logger) sourced() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(2)
	fn := "(unknown)"
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
		fn = runtime.FuncForPC(pc).Name()
	}

	fnS := strings.Split(fn, "/")[rootPackageSlash:]
	fn = strings.Join(fnS[:], "/")
	logger := l.entry.WithField("source", fmt.Sprintf("%s/%s/%s:%d", workspace, strings.Split(fn, ".")[0], file, line))
	return logger
}

var origLogger = logrus.New()
var baseLogger = logger{entry: logrus.NewEntry(origLogger)}

// New returns a new logger.
func New() Logger {
	return logger{entry: logrus.NewEntry(origLogger)}
}

// Base returns the base logger.
func Base() Logger {
	return baseLogger
}

//SetWorkspace sets the workspace of the logger
func (l logger) SetWorkspace(path string) {
	workspace = path
}

//SetRootPackageSlash sets the fileSlash of the logger
func (l logger) SetRootPackageSlash(number int) {
	rootPackageSlash = number
}

// SetLevel sets the Level of the base logger
func SetLevel(level Level) {
	baseLogger.entry.Logger.Level = logrus.Level(level)
}

// SetOut sets the output destination base logger
func SetOut(out io.Writer) {
	baseLogger.entry.Logger.Out = out
}

// With attaches a key,value pair to a logger.
func With(key string, value interface{}) Logger {
	return baseLogger.With(key, value)
}

// WithError returns a Logger that will print an error along with the next message.
func WithError(err error) Logger {
	return logger{entry: baseLogger.sourced().WithError(err)}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	baseLogger.sourced().Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	baseLogger.sourced().Debugln(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	baseLogger.sourced().Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	baseLogger.sourced().Infoln(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	baseLogger.sourced().Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	baseLogger.sourced().Warnln(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	baseLogger.sourced().Error(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	baseLogger.sourced().Errorln(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	baseLogger.sourced().Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	baseLogger.sourced().Fatalln(args...)
}
