package logger

import "github.com/dustin/go-humanize"
import "github.com/segmentio/go-log"
import "net/http"
import "time"

// Logger middleware.
type Logger struct {
	h http.Handler
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
		return &Logger{h}
	}
}

// ServeHTTP implementation.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := &wrapper{w, 0, 200}
	log.Info(">> %s %s", r.Method, r.RequestURI)
	l.h.ServeHTTP(res, r)
	size := humanize.Bytes(uint64(res.written))
	if res.status >= 500 {
		log.Error("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	} else {
		log.Info("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	}
}
