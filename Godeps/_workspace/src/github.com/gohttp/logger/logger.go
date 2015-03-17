package logger

import "github.com/dustin/go-humanize"
import "github.com/segmentio/go-log"
import "net/http"
import "time"

// Logger middleware.
type Logger struct {
	h   http.Handler
	log *log.Logger
}

// SetLogger sets the logger to `log`.
func (l *Logger) SetLogger(log *log.Logger) {
	l.log = log
}

// wrapper to capture status.
type wrapper struct {
	http.ResponseWriter
	written int
	status  int
}

// capture status.
func (w *wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// capture written bytes.
func (w *wrapper) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.written += n
	return n, err
}

// New logger middleware.
func New() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &Logger{
			log: log.Log,
			h:   h,
		}
	}
}

// NewLogger logger middleware with the given logger.
func NewLogger(log *log.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &Logger{
			log: log,
			h:   h,
		}
	}
}

// ServeHTTP implementation.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := &wrapper{w, 0, 200}

	l.log.Info(">> %s %s", r.Method, r.RequestURI)
	l.h.ServeHTTP(res, r)
	size := humanize.Bytes(uint64(res.written))

	switch {
	case res.status >= 500:
		l.log.Error("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	case res.status >= 400:
		l.log.Warning("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	default:
		l.log.Info("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	}
}
