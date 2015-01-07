package serve

import "github.com/segmentio/go-log"
import "path/filepath"
import "net/http"
import "path"

// New file serving middleware, restricted to `dir`.
func New(dir string) func(http.Handler) http.Handler {
	log := log.Log.New("serve " + dir)
	fs := http.Dir(dir)

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Path

			log.Debug("open %s", name)

			f, err := fs.Open(name)

			if err != nil {
				log.Debug("%s open error: %s", name, err)
				h.ServeHTTP(w, r)
				return
			}

			s, err := f.Stat()

			if err != nil {
				log.Debug("%s stat error: %s", name, err)
				h.ServeHTTP(w, r)
				return
			}

			name = filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name)))

			log.Debug("serving %s (%d bytes)", name, s.Size)
			http.ServeFile(w, r, name)
		})
	}
}
