package quest

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeARequest(t *testing.T) {
	q, _ := Request(GET, "http://httpbin.org/get")
	Convey("Should be making a Request", t, func() {
		So(q.Method, ShouldEqual, GET)
	})
}

func TestDownload(t *testing.T) {
	os.Mkdir("tmp", os.ModePerm)

	Convey("Downloading file", t, func() {
		Convey("Downloading stream.log in progress\n", func() {
			q, _ := Download(GET, "http://httpbin.org/bytes/1024", "tmp/stream.log")
			q.
				Progress(func(c, t, e int64) {
				log.Println(c, t, e)
				So(c, ShouldBeLessThanOrEqualTo, t)
			}).Do()
		})
		Convey("Downloading stream2.log in progress and invoke response handler\n", func() {
			var n int64
			stream2, _ := os.Create("tmp/stream2.log")
			q, _ := Download(GET, "http://httpbin.org/bytes/10240", stream2)
			q.
				Progress(func(c, t, e int64) {
				n = c
				log.Println(c, t, e)
				So(c, ShouldBeLessThanOrEqualTo, t)
			}).Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
				l := int64(data.Len())
				So(n, ShouldEqual, l)
				So(res.ContentLength, ShouldEqual, l)
			})
		})
	})
}

func TestUpload(t *testing.T) {
	Convey("Uploading file", t, func(m C) {
		m.Convey("Uploading one file\n", func() {
			data := map[string]interface{}{
				"stream": "tmp/stream.log",
			}
			q, _ := Upload(POST, "http://httpbin.org/post", data)
			q.
				Progress(func(c, t, e int64) {
				log.Println(c, t, e)
				m.So(c, ShouldBeLessThanOrEqualTo, t)
			}).Do()
		})
		Convey("Uploading multi files\n", func() {
			stream2, _ := os.Open("quest_test.go")
			stream3 := bytes.NewBufferString(`Hello Quest!`)
			data := map[string]interface{}{
				"stream1": "quest.go", // filepath or filename
				"stream2": stream2,    // *os.File
				"stream3": stream3,    // io.Reader, filename is fieldname `stream3`
			}

			q, _ := Upload(POST, "http://httpbin.org/post", data)
			q.
				Parameters(map[string]string{"foo": "bar", "bar": "foo"}).
				Progress(func(c, t, e int64) {
				log.Println(c, t, e)
				m.So(c, ShouldBeLessThanOrEqualTo, t)
			}).Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
				l := int64(data.Len())
				m.So(res.ContentLength, ShouldEqual, l)
			})
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Get Request\n", t, func() {
		q, _ := Get("http://httpbin.org/get")
		q.Do()
		So(q.res.StatusCode, ShouldEqual, 200)
	})
}

func TestPost(t *testing.T) {
	Convey("Post Request\n", t, func() {
		q, _ := Post("http://httpbin.org/post")
		q.Do()
		So(q.res.StatusCode, ShouldEqual, 200)
	})
}

func TestPatch(t *testing.T) {
	Convey("Patch Request\n", t, func() {
		q, _ := Patch("http://httpbin.org/patch")
		q.Do()
		So(q.res.StatusCode, ShouldEqual, 200)
	})
}

func TestPut(t *testing.T) {
	Convey("Put Request\n", t, func() {
		q, _ := Put("http://httpbin.org/put")
		q.Do()
		So(q.res.StatusCode, ShouldEqual, 200)
	})
}

func TestDelete(t *testing.T) {
	Convey("Delete Request\n", t, func() {
		q, _ := Put("http://httpbin.org/delete")
		q.Do()
		So(q.res.StatusCode, ShouldEqual, 405)
	})
}
