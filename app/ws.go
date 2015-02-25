package app

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

const PREFIX = "/ws"

func WSHandler(prefix string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		if prefix == "" {
			prefix = PREFIX
		}
		ws := sockjs.NewHandler(prefix, sockjs.DefaultOptions, echoHandler)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, prefix) {
				ws.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func echoHandler(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			fmt.Println(msg)
			session.Send(msg)
			continue
		}
		break
	}
}
