package dashboard

import (
	"fmt"
	"net/http"
	"path"
	"runtime"

	"github.com/codegangsta/negroni"
	hr "github.com/julienschmidt/httprouter"

	// middlewares
	middle "github.com/dockerboard/dockerboard/server/middleware"
)

func ApiIndex(w http.ResponseWriter, r *http.Request, _ hr.Params) {
	fmt.Fprint(w, "Coming soon!\n")
}

func Serve() {
	_, filename, _, _ := runtime.Caller(0)
	// Note config dir
	dir := path.Join(path.Dir(filename), "../client/src")

	router := hr.New()
	router.GET("/api/v1", ApiIndex)

	n := negroni.Classic()
	n.Use(middle.Logrus())
	n.Use(negroni.NewStatic(http.Dir(dir)))
	n.UseHandler(router)
	n.Run(":3333")
}
