package dashboard

import (
	"net/http"
	"path"
	"runtime"

	"github.com/codegangsta/negroni"
	hr "github.com/julienschmidt/httprouter"

	// middlewares
	middle "github.com/dockerboard/dockerboard/server/middleware"
)

func Serve() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../client")
	router := hr.New()
	router.ServeFiles("/*filepath", http.Dir(dir))
	n := negroni.Classic()
	n.Use(middle.Logrus())
	n.UseHandler(router)
	n.Run(":3333")
}
