package response

import "net/http"

// Error responds with a generic status code response.
func Error(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
