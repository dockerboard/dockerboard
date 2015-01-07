# quest

[![build status][travis-image]][travis-url]
[![GoDoc][godoc-image]][godoc-url]


**quest** is an elegant HTTP networking library written in Go. [Alamofire][] inspired.


## Features

- [x] Chainable Request / Response methods
- [x] URL / JSON / Parameter Encoding
- [x] Progress Tracking
- [x] Upload File / Data / Stream
- [x] Download using Request
- [x] Authentication
- [x] Timeout, defaults to 30 * time.Second
- [x] cURL Debug Output
- [x] TLS config
- [x] Support Short APIS `Get`, `Post`, `Patch`, `Put`, `Delete`
- [ ] Pipe Stream to other request
- [ ] Download Resume data
- [ ] More Errors Output
- [ ] HTTP Response Validation
- [ ] Comprehensive Unit Test Coverage
- [ ] Complete Documentation



## Usage


### Making a Request

```go
import "github.com/go-libs/quest"

q, err := quest.Request("GET", "http://httpbin.org/get")
```


### Response Handling

```go
q, err := quest.Request(quest.GET, "http://httpbin.org/get")
q.
  Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, e error) {
  log.Println(req)
  log.Println(res)
  log.Println(data)
  log.Println(err)
})
```


### Response Serialization

Built-in Response Methods

* `Response(func(*http.Request, *http.Response, *bytes.Buffer, error))`
* `ResponseBytes(func(*http.Request, *http.Response, []byte, error))`
* `ResponseString(func(*http.Request, *http.Response, string, error))`
* `ResponseJSON(f interface{})`, `f` ___Must___ be `func`
    - `func(req *http.Request, res *http.Response, data *JSONStruct, e error)`
    - `func(req *http.Request, res *http.Response, data JSONStruct, e error)`


#### Response String Handler

```go
q, _ := quest.Request(quest.GET, "http://httpbin.org/get")
q.ResponseString(func(req *http.Request, res *http.Response, data string, e error) {
  log.Println(data)
})
```


#### Response JSON Handler

```go
type DataStruct struct {
  Headers map[string]string
  Origin  string
}

q, _ := quest.Request(quest.GET, "http://httpbin.org/get")
q.ResponseJSON(func(req *http.Request, res *http.Response, data DataStruct, e error) {
  log.Println(data)
})

q.ResponseJSON(func(req *http.Request, res *http.Response, data *DataStruct, e error) {
  log.Println(data)
})
```


#### Chained Response Handlers

Response handlers can even be chained:
```go
q, _ := quest.Request(quest.GET, "http://httpbin.org/get")
q.
  ResponseString(func(req *http.Request, res *http.Response, data string, e error) {
  log.Println(data)
}).
  ResponseJSON(func(req *http.Request, res *http.Response, data *DataStruct, e error) {
  log.Println(data)
})
```


### HTTP Methods

```
OPTIONS
GET
HEAD
POST
PUT
PATCH
DELETE
TRACE
CONNECT
```


### Query String

```go
type Options struct {
  Foo string `url:"foo"`
}

q, _ := quest.Request(quest.GET, "http://httpbin.org/get")
q.Query(Options{"bar"})
// http://httpbin.org/get?foo=bar
```


### POST Request with JSON-encoded Parameters

```go
type PostParameters struct {
  Foo []int             `json:"foo,omitempty"`
  Bar map[string]string `json:"bar,omitempty"`
}

parameters := PostParameters{
  "foo": []int{1, 2, 3},
  "bar": map[string]string{"baz": "qux"},
}

type DataStruct struct {
  Headers map[string]string
  Origin  string
  Json    PostParameters `json:"json,omitempty"`
}

type OtherDataStruct struct {
  Headers map[string]string
  Origin  string
}

q, _ := quest.Request(quest.POST, "http://httpbin.org/post")
q.Encoding("JSON").
  Parameters(&parameters).
  ResponseJSON(func(req *http.Request, res *http.Response, data *DataStruct, e error) {
  log.Println(data)
}).
  ResponseJSON(func(req *http.Request, res *http.Response, data OtherDataStruct, e error) {
  log.Println(data)
})
```



### Downloading


#### Downloading a File

```go
q, _ := quest.Download(quest.GET, "http://httpbin.org/stream/100", "stream.log")
q.Do()
```


#### Downloading a File w/Progress

```go
destination := "tmp/stream.log"
q, _ := quest.Download(quest.GET, "http://httpbin.org/bytes/1024", destination)
q.
  Progress(func(bytesRead, totalBytesRead, totalBytesExpectedToRead int64) {
    log.Println(bytesRead, totalBytesRead, totalBytesExpectedToRead)
  }).Do()

destination, _ := os.Create("tmp/stream2.log")
q, _ := quest.Download(quest.GET, "http://httpbin.org/bytes/10240", destination)
q.
  Progress(func(current, total, expected int64) {
    log.Println(current, total, expected)
  }).Response(func(request *http.Request, response *http.Response, data *bytes.Buffer, err error) {
    log.Println(data.Len())
  })
```


### Uploading


#### Supported Upload Types

* File
* Data
* Stream


#### Uploading a File

```go
q, _: = quest.Upload(quest.POST, "http://httpbin.org/post", map[string]string{"stream": "tmp/stream.log"})
q.Do()
```


#### Uploading multi files w/Progress

```go
stream2, _ := os.Open("tmp/stream2.log")
stream3 := bytes.NewBufferString(`Hello Quest!`)
data := map[string]interface{}{
  "stream1": "tmp/stream.log",      // filepath or filename
  "stream2": stream2,               // *os.File
  "stream3": stream3,               // io.Reader, filename is fieldname `stream3`
}

q, _ := quest.Upload(quest.POST, "http://httpbin.org/post", data)
q.
  Progress(func(current, total, expected int64) {
    log.Println(current, total, expected)
  }).Response(func(req *http.Request, res *http.Response, data *bytes.Buffer, err error) {
    log.Println(data.Len())
  })
```


### Authenticate

#### HTTP Basic Authentication

```go
type Auth struct {
  User          string
  Passwd        string
  Authenticated bool
}
user := "user"
passwd := "password"

q, _ := quest.Request(quest.GET, "https://httpbin.org/basic-auth/"+user+"/"+passwd)
q.Authenticate(user, passwd).
  ResponseJSON(func(_ *http.Request, _ *http.Response, data Auth, _ error) {
  log.Println(data)
}).Do()
```


### Timeout

```go
s := time.Duration(3 * time.Second)
q, _ := quest.Request(quest.GET, "https://httpbin.org/delay/5")
q.Timeout(s).Do()
```


### Println & DebugPrintln Request

```go
q, _ := Request(GET, "http://httpbin.org/cookies")
log.Println(q.Println())
```

```go
c := &http.Cookie{}
c.Name = "k"
c.Value = "v"
q, _ := Request(GET, "http://httpbin.org/cookies")
q.Query(&queryParams)
q.Cookie(c)
log.Println(q.DebugPrintln())
```


### Response `Bytes()`, `String()`, `JSON(v interface{})`

```go
q, _ := Request(GET, "http://httpbin.org/get")
b1, err := q.Bytes()
s1, err := q.String()

type DataStruct struct {
  Headers map[string]string
  Origin  string
}

v := DataStruct{}
err := q.JSON(&v)
```


### TLS client config `TLSConfig(*tls.Config)`

```go
t := &tls.Config{}
q, _ := Request(GET, "http://httpbin.org/get")
q.TLSConfig(t)
```


### `Get()` `Post()` `Put()` `Patch()` `Delete()` Short APIs

```go
q, _ = Get("http://httpbin.org/get")
```


## License

MIT

[Alamofire]: https://github.com/Alamofire/Alamofire
[travis-image]: https://img.shields.io/travis/go-libs/quest/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/go-libs/quest
[godoc-image]: https://godoc.org/github.com/go-libs/quest?status.svg
[godoc-url]: http://godoc.org/github.com/go-libs/quest
