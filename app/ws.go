package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/igm/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

// Websocket prefix.
const PREFIX = "/ws"

type WSDataSchema struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

var chat pubsub.Publisher

// WSHandler a middleware for `gohttp`.
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
	log.Println("new sockjs session established")
	var closedSession = make(chan struct{})
	session.Send(session.ID())
	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					return
				}
			}

		}
	}()
	for {
		if msg, err := session.Recv(); err == nil {
			data := &WSDataSchema{}
			buf := bytes.NewBufferString(msg)
			err = json.Unmarshal(buf.Bytes(), data)
			if err == nil {
				_type := data.Type
				_endpoint := data.Endpoint
				_id := data.Id
				if _type == "container" {
					session.Send(_endpoint + _id)
					break
				}
			}
			chat.Publish(msg)
			//session.Send(msg)
			continue
		}
		break
	}
	close(closedSession)
	log.Println("sockjs session closed")
}
