//
// Default logger:
//
//   log.Info("something %s", "here")
//
package log

import "os"

var Log = New(os.Stderr, INFO, "")

// SetPrefix wrapper.
func SetPrefix(str string) {
	Log.SetPrefix(str)
}

// SetLevel wrapper.
func SetLevel(level Level) {
	Log.SetLevel(level)
}

// SetLevelString wrapper.
func SetLevelString(level string) {
	Log.SetLevelString(level)
}

// Debug log.
func Debug(msg string, args ...interface{}) error {
	return Log.Debug(msg, args...)
}

// Info log.
func Info(msg string, args ...interface{}) error {
	return Log.Info(msg, args...)
}

// Notice log.
func Notice(msg string, args ...interface{}) error {
	return Log.Notice(msg, args...)
}

// Warning log.
func Warning(msg string, args ...interface{}) error {
	return Log.Warning(msg, args...)
}

// Error log.
func Error(msg string, args ...interface{}) error {
	return Log.Error(msg, args...)
}

// Critical log.
func Critical(msg string, args ...interface{}) error {
	return Log.Critical(msg, args...)
}

// Alert log.
func Alert(msg string, args ...interface{}) error {
	return Log.Alert(msg, args...)
}

// Emergency log.
func Emergency(msg string, args ...interface{}) error {
	return Log.Emergency(msg, args...)
}

// Check if there's an `err` and exit, useful for bootstrapping.
func Check(err error) {
	Log.Check(err)
}

// Fatalf is equivalent to Error() followed by a call to os.Exit(1).
func Fatalf(msg string, args ...interface{}) {
	Log.Fatalf(msg, args...)
}
