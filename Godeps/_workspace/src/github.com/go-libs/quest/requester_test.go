package quest

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponseHandling(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Set("foo", "bar")
	queryParams.Set("name", "活力")

	Convey("Query, query string", t, func() {
		q, _ := Request(GET, "http://httpbin.org/get")
		q.
			Query(&queryParams).
			Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
			So(req.URL.String(), ShouldEqual, "http://httpbin.org/get?foo=bar&name=%E6%B4%BB%E5%8A%9B")
		})
	})

	Convey("Parameters, ContentLength should equal to buffer length", t, func() {
		q, _ := Request(POST, "http://httpbin.org/post")
		q.
			Parameters(queryParams).
			Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
			So(res.ContentLength, ShouldEqual, int64(data.Len()))
		})
	})

	parameters := map[string]interface{}{
		"foo": []int{1, 2, 3},
		"bar": map[string]string{"baz": "qux"},
	}

	type DataStruct struct {
		Headers map[string]string
		Origin  string
	}

	type DataStruct2 struct {
		Origin string
	}

	Convey("Response JSON", t, func() {
		q, _ := Request(POST, "http://httpbin.org/post")
		q.
			Encoding("JSON").
			Parameters(parameters).
			ResponseJSON(func(req *http.Request, res *http.Response, data *DataStruct, e error) {
			Convey("Data - a pointer struct", func() {
				So(data, ShouldPointTo, data)
				So(data.Headers["Host"], ShouldEqual, "httpbin.org")
			})
		}).
			ResponseJSON(func(req *http.Request, res *http.Response, data DataStruct2, e error) {
			Convey("Data - a struct", func() {
				So(&data, ShouldNotPointTo, &DataStruct2{})
				So(data.Origin, ShouldNotBeNil)
			})
		}).
			// Nothing happend
			ResponseJSON(func(req *http.Request, res *http.Response, data DataStruct2, e error, g error) {
			log.Println("Nothing happend!")
		})
	})

	type PostParameters struct {
		Foo []int             `json:"foo,omitempty"`
		Bar map[string]string `json:"bar,omitempty"`
	}

	parameters2 := &PostParameters{
		Foo: []int{1, 2, 3},
		Bar: map[string]string{"baz": "qux"},
	}

	type DataStruct4 struct {
		Origin string
	}
	type DataStruct3 struct {
		Headers map[string]string
		Origin  string
		Json    PostParameters `json:"json,omitempty"`
	}

	Convey("Response JSON, using JSON decode", t, func() {
		q, _ := Request(POST, "http://httpbin.org/post")
		q.
			Encoding("JSON").
			Parameters(parameters2).
			ResponseJSON(func(req *http.Request, res *http.Response, data DataStruct4, e error) {
			Convey("Using DataStruct4 JSON struct", func() {
				So(data.Origin, ShouldNotBeNil)
			})
		}).
			ResponseJSON(func(req *http.Request, res *http.Response, data *DataStruct3, e error) {
			Convey("Using DataStruct3 JSON struct", func() {
				So(&data.Json, ShouldResemble, parameters2)
			})
		})
	})

	Convey("Encoding Query Options", t, func() {
		type Options struct {
			Foo string `url:"foo"`
			Baz []int  `url:"baz"`
		}

		// http://httpbin.org/get
		q, _ := Request(GET, "http://httpbin.org/get")
		q.
			Query(Options{"bar", []int{233, 377, 610}}).
			Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
			So(req.URL.String(), ShouldEqual, "http://httpbin.org/get?baz=233&baz=377&baz=610&foo=bar")
		})
	})
}

func TestPrintln(t *testing.T) {
	q, _ := Request(GET, "http://httpbin.org/get")
	Convey("Println request", t, func() {
		So(q.Println(), ShouldEqual, "GET http://httpbin.org/get")
	})
	q, _ = Request(GET, "http://httpbin.org/get")
	q.Do()
	Convey("Println request", t, func() {
		So(q.Println(), ShouldEqual, "GET http://httpbin.org/get "+strconv.Itoa(q.res.StatusCode))
	})
}

func TestDebugPrintln(t *testing.T) {
	c1 := &http.Cookie{}
	c1.Name = "k1"
	c1.Value = "v1"
	c2 := &http.Cookie{}
	c2.Name = "k2"
	c2.Value = "v2"
	queryParams := url.Values{}
	queryParams.Set("foo", "bar")
	queryParams.Set("name", "bazz")
	q, _ := Request(GET, "http://httpbin.org/cookies")
	q.Query(&queryParams)
	q.Cookie(c1, c2)
	Convey("DebugPrintln request", t, func() {
		s := []string{"$ curl -i", "-b " + strconv.Quote("k1=v1; k2=v2"), strconv.Quote("http://httpbin.org/cookies?foo=bar&name=bazz")}
		So(q.DebugPrintln(), ShouldEqual, strings.Join(s, " \\\n\t"))
	})
}

func TestAuthenticate(t *testing.T) {
	type Auth struct {
		User          string
		Passwd        string
		Authenticated bool
	}
	user := "user"
	passwd := "password"

	Convey("Authenticate", t, func() {
		Convey("Basic Auth", func() {
			q, _ := Request(GET, "https://httpbin.org/basic-auth/"+user+"/"+passwd)
			q.Authenticate(user, passwd).
				ResponseJSON(func(_ *http.Request, _ *http.Response, data Auth, _ error) {
				So(data.User, ShouldEqual, user)
				So(data.Authenticated, ShouldEqual, true)
			}).Do()
		})
	})
}

func TestTimeout(t *testing.T) {
	Convey("Timeout", t, func() {
		Convey("It's timeout.", func() {
			s := time.Duration(3 * time.Second)
			q, _ := Request(GET, "https://httpbin.org/delay/5")
			q.Timeout(s).Do()
		})
		Convey("It's not timeout.", func() {
			s := time.Duration(30 * time.Second)
			q, _ := Request(GET, "https://httpbin.org/delay/5")
			q.Timeout(s).Do()
		})
	})
}

func TestSetHeader(t *testing.T) {
	type DataStruct struct {
		Headers map[string]string
	}
	Convey("set header", t, func() {
		q, _ := Request("GET", "http://httpbin.org/headers")
		q.Set("Quest", "Test").
			ResponseJSON(func(_ *http.Request, _ *http.Response, data DataStruct, _ error) {
			So(data.Headers["Quest"], ShouldEqual, "Test")
		}).Do()
	})
}

func TestBytesNotHandler(t *testing.T) {
	queryParams := url.Values{}
	queryParams.Set("foo", "bar")
	queryParams.Set("name", "活力")

	Convey("Response Bytes not handler", t, func() {
		q, _ := Request(GET, "http://httpbin.org/get")
		_, err := q.Query(&queryParams).Bytes()
		So(err, ShouldBeNil)
	})
}

func TestStringNotHandler(t *testing.T) {
	Convey("Response String not handler", t, func() {
		q, _ := Request(GET, "http://httpbin.org/get")
		_, err := q.String()
		So(err, ShouldBeNil)
	})
}

func TestJSONNotHandler(t *testing.T) {
	type DataStruct struct {
		Headers map[string]string
		Origin  string
	}
	Convey("Response JSON not handler", t, func() {
		q, _ := Request(GET, "http://httpbin.org/get")
		b := DataStruct{}
		err := q.JSON(&b)
		So(err, ShouldBeNil)
		So(b.Headers["Host"], ShouldEqual, "httpbin.org")
	})
}
