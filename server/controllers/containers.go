package controllers

import (
	"fmt"
	"net/http"
)

type ContainersController struct{}

func (cc *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "docker containers!\n")
}
