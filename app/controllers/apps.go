package controllers

import (
	"fmt"
	"net/http"
)

type AppsController struct{}

func (a *AppsController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "docker apps!\n")
}