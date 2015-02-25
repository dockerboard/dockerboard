package app

import (
	"fmt"
	"net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Using docker@v1.17 API.")
}
