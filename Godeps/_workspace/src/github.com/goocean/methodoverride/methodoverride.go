package methodoverride

import "net/http"

const (
	HEAD_X_HTTP_METHOD_OVERRIDE = "X-HTTP-Method-Override"
	PARAM_HTTP_METHOD_OVERRIDE  = "_method"
)

var Methods = []string{"PUT", "PATCH", "DELETE"}

func New() func(http.Handler) http.Handler {
	return MethodOverride
}

// Check methods
func supports(m string) bool {
	for _, v := range Methods {
		if v == m {
			return true
		}
	}
	return false
}

// Get `_method` From Head Or Form
func getMethod(r *http.Request) (m string) {
	m = r.Header.Get(HEAD_X_HTTP_METHOD_OVERRIDE)
	if m == "" {
		m = r.FormValue(PARAM_HTTP_METHOD_OVERRIDE)
	}
	return
}

// Method Override Handler
func MethodOverride(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			m := getMethod(r)
			if m != "" {
				if supports(m) {
					r.Method = m
				}
			}
		}

		h.ServeHTTP(w, r)
	})
}
