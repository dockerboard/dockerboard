package app

import (
	"fmt"
	"net/http"
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Coming soon!\n")
}