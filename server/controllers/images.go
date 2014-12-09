package controllers

import (
	"fmt"
	"net/http"
)

type ImagesController struct{}

func (ic *ImagesController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "docker images!\n")
}
