//
// Simple logger similar to Go's standard logger with log level.
//
//   l := log.New(os.Stderr, INFO, "myapp")
//   l.Error("something exploded")
//
package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Level.
type Level int

// Levels.
const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
	ALERT
	EMERGENCY
)

// Level map.
var levels = map[string]Level{
	"debug":     DEBUG,
	"info":      INFO,
	"notice":    NOTICE,
	"warning":   WARNING,
	"error":     ERROR,
	"critical":  CRITICAL,
	"alert":     ALERT,
	"emergency": EMERGENCY,
}

// Logger.
type Logger struct {
	Writer io.Writer
	Level  Level
	Prefix string
	sync.Mutex
}

// New logger which writes to `w` at the given `level`. Optionally
// provide a `prefix` for the logger.
func New(w io.Writer, level Level, prefix string) *Logger {
	l := &Logger{Writer: w, Level: level, Prefix: prefix}
	l.SetPrefix(prefix)
	l.SetLevelFromEnv("LOG_LEVEL")
	return l
}

// SetLevelFromEnv forces the log level based on the given
// environment variable `name` when present.
func (l *Logger) SetLevelFromEnv(name string) {
	if s := os.Getenv("LOG_LEVEL"); s != "" {
		l.SetLevel(levels[s])
	}
}

// SetPrefix changes the prefix to `str`.
func (l *Logger) SetPrefix(str string) {
	l.Lock()
	defer l.Unlock()

	if str != "" {
		str = " " + str + ":"
	}

	l.Prefix = str
}

// New logger which inherits the writer and level.
func (l *Logger) New(prefix string) *Logger {
	return New(l.Writer, l.Level, prefix)
}

// SetLevel changes the log `level`.
func (l *Logger) SetLevel(level Level) {
	l.Lock()
	defer l.Unlock()

	l.Level = level
}

// Write to the logger.
func (l *Logger) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b, []byte("\n"))
	for _, line := range lines {
		l.Info("%s", string(line))
	}
	return len(b), nil
}

// SetLevelString changes the log `level` via string.
// This is especially useful for providing a command-line
// flag to your program to adjust the level.
//
// If the level string is invalid an error is returned.
func (l *Logger) SetLevelString(level string) error {
	l.Lock()
	defer l.Unlock()

	if val, ok := levels[level]; ok {
		l.Level = val
		return nil
	}

	return fmt.Errorf("%q is not a valid level", level)
}

// Log a message.
func (l *Logger) Log(lvl string, level Level, msg string, args ...interface{}) error {
	l.Lock()
	defer l.Unlock()

	if l.Level > level {
		return nil
	}

	ts := time.Now().Format("2006-01-02 15:04:05")
	f := fmt.Sprintf("%s [%s]%s %s\n", ts, lvl, l.Prefix, msg)
	_, err := fmt.Fprintf(l.Writer, f, args...)
	return err
}

// Debug log.
func (l *Logger) Debug(msg string, args ...interface{}) error {
	return l.Log("DEBUG", DEBUG, msg, args...)
}

// Info log.
func (l *Logger) Info(msg string, args ...interface{}) error {
	return l.Log("INFO", INFO, msg, args...)
}

// Notice log.
func (l *Logger) Notice(msg string, args ...interface{}) error {
	return l.Log("NOTICE", NOTICE, msg, args...)
}

// Warning log.
func (l *Logger) Warning(msg string, args ...interface{}) error {
	return l.Log("WARNING", WARNING, msg, args...)
}

// Error log.
func (l *Logger) Error(msg string, args ...interface{}) error {
	return l.Log("ERROR", ERROR, msg, args...)
}

// Critical log.
func (l *Logger) Critical(msg string, args ...interface{}) error {
	return l.Log("CRITICAL", CRITICAL, msg, args...)
}

// Alert log.
func (l *Logger) Alert(msg string, args ...interface{}) error {
	return l.Log("ALERT", ALERT, msg, args...)
}

// Emergency log.
func (l *Logger) Emergency(msg string, args ...interface{}) error {
	return l.Log("EMERGENCY", EMERGENCY, msg, args...)
}

// Fatalf is equivalent to Error() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.Check(fmt.Errorf(msg, args...))
}

// Check if there's an `err` and exit, useful for bootstrapping.
func (l *Logger) Check(err error) {
	if err != nil {
		Log.Error("exiting: %s", err.Error())
		os.Exit(1)
	}
}
