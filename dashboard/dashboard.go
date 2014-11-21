package dashboard

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	hr "github.com/julienschmidt/httprouter"

	// middlewares
	middle "github.com/dockerboard/dockerboard/dashboard/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ hr.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Serve() {
	router := hr.New()
	router.GET("/", Index)

	n := negroni.Classic()
	n.Use(middle.Logrus())
	n.UseHandler(router)
	n.Run(":3333")
}
