package response

import "encoding/json"
import "net/http"

// JSON response with optional status code.
func JSON(w http.ResponseWriter, val interface{}, code ...int) {
	var b []byte
	var err error

	if Pretty {
		b, err = json.MarshalIndent(val, "", "  ")
	} else {
		b, err = json.Marshal(val)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	w.Write(b)
}
