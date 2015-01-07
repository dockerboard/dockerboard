package quest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-libs/progress"
	"github.com/go-libs/syncreader"
)

var defaultTimeout = 30 * time.Second

type HandlerFunc func(*http.Request, *http.Response, *bytes.Buffer, error)
type BytesHandlerFunc func(*http.Request, *http.Response, []byte, error)
type StringHandlerFunc func(*http.Request, *http.Response, string, error)

// A Request manages communication with http service.
type Requester struct {
	// HTTP method
	Method string

	// Base URL for Requests.
	Endpoint string
	Url      *url.URL

	// HTTP client
	client *http.Client

	// Other client, unix
	clientConn *httputil.ClientConn

	// HTTP request
	req *http.Request

	Header  http.Header
	Body    io.ReadCloser
	Length  int64
	rawBody interface{}

	// HTTP response
	res *http.Response

	StatusCode int

	// response body, buffer
	isBodyClosed bool
	Buffer       *bytes.Buffer

	err error

	// Upload
	IsUpload bool
	files    map[string]interface{}

	// Download
	IsDownload  bool
	destination interface{}

	// Cookies
	cookies []*http.Cookie

	// Progress
	pg *progress.Progress

	transport *http.Transport
	tlsconfig *tls.Config
	timeout   time.Duration
}

func (r *Requester) Files(files map[string]interface{}) *Requester {
	r.files = files
	return r
}

func (r *Requester) Destination(destination interface{}) *Requester {
	r.destination = destination
	return r
}

func (r *Requester) Timeout(t time.Duration) *Requester {
	r.timeout = t
	return r
}

func (r *Requester) Set(key, value string) *Requester {
	r.Header.Set(key, value)
	return r
}

func (r *Requester) Query(data interface{}) *Requester {
	qs, err := QueryString(data)
	r.err = err
	r.Url.RawQuery = qs
	return r
}

func (r *Requester) Parameters(data interface{}) *Requester {
	if encodesParametersInURL(r.Method) {
		r.err = errors.New("Must not be GET, HEAD, DELETE methods.")
		return r
	}
	r.rawBody = data
	return r
}

func (r *Requester) packBody() {
	if r.rawBody == nil {
		return
	}
	body, length, err := packBody(r.rawBody)
	r.err = err
	if length > 0 && body != nil {
		r.Body = body
		r.Length = length
	}
}

func (r *Requester) Form(files map[string]interface{}, fields map[string]string) *Requester {
	var data interface{}
	if len(files) > 0 {
		pr, pw := io.Pipe()
		mw := multipart.NewWriter(pw)
		go func() {
			var (
				fp io.Writer
				fr io.Reader
			)
			for fieldname, file := range files {
				fh, name, err := getFile(file)
				if err == nil {
					file = fh
				} else {
					name = fieldname
				}
				fp, err = mw.CreateFormFile(fieldname, filepath.Base(name))
				if err != nil {
					log.Fatal(err)
				}
				fr, _ = file.(io.Reader)
				_, err = io.Copy(fp, ioutil.NopCloser(fr))
				if err != nil {
					log.Fatal(err)
				}
			}

			for k, v := range fields {
				mw.WriteField(k, v)
			}
			mw.Close()
			pw.Close()
		}()
		r.Header.Set("Content-Type", mw.FormDataContentType())
		data = pr
	} else {
		data = fields
	}
	r.rawBody = data
	return r
}

func (r *Requester) Encoding(t string) *Requester {
	t = strings.ToUpper(t)
	if t == "JSON" {
		t = "application/json"
	}
	if t != "" {
		r.Header.Set("Content-Type", t)
	}
	return r
}

func (r *Requester) Authenticate(username, password string) *Requester {
	r.Url.User = url.UserPassword(username, password)
	return r
}

func (r *Requester) Cookie(cookies ...*http.Cookie) *Requester {
	r.cookies = cookies
	return r
}

func (r *Requester) Progress(f progress.HandlerFunc) *Requester {
	r.pg = progress.New()
	r.pg.Progress = f
	return r
}

func (r *Requester) response() (*bytes.Buffer, error) {
	if r.err != nil {
		return r.Buffer, r.err
	}
	if r.isBodyClosed {
		return r.Buffer, nil
	}
	r.isBodyClosed = true
	return r.Do()
}

func (r *Requester) Response(handler HandlerFunc) *Requester {
	_, err := r.response()
	handler(r.req, r.res, r.Buffer, err)
	return r
}

func (r *Requester) ResponseBytes(handler BytesHandlerFunc) *Requester {
	_, err := r.response()
	handler(r.req, r.res, r.Buffer.Bytes(), err)
	return r
}

func (r *Requester) ResponseString(handler StringHandlerFunc) *Requester {
	_, err := r.response()
	handler(r.req, r.res, r.Buffer.String(), err)
	return r
}

func (r *Requester) ResponseJSON(f interface{}) *Requester {
	var (
		fn                    = reflect.ValueOf(f)
		t                     = reflect.TypeOf(f)
		argsNum               = t.NumIn()
		in                    = make([]reflect.Value, argsNum) //Panic if t is not kind of Func
		reqV, resV, dataV, eV reflect.Value
		err                   error
	)
	if argsNum != 4 {
		err = errors.New("ResponseJSON: invalid arguments.")
		return r
	} else {
		_, err = r.response()
		if err != nil {
			dataV = reflect.New(t.In(2)).Elem()
		} else {
			dataT := t.In(2)
			dataK := dataT.Kind()
			if dataK == reflect.Ptr {
				dataT = dataT.Elem()
			}
			dataN := reflect.New(dataT)
			data := dataN.Interface()
			err = json.Unmarshal(r.Buffer.Bytes(), &data)
			dataV = reflect.ValueOf(data)
			if dataK != reflect.Ptr {
				dataV = reflect.Indirect(dataV)
			}
		}
	}
	if err == nil {
		eV = reflect.New(t.In(3)).Elem()
	} else {
		eV = reflect.ValueOf(err)
	}
	reqV = reflect.ValueOf(r.req)
	resV = reflect.ValueOf(r.res)
	in[0] = reqV
	in[1] = resV
	in[2] = dataV
	in[3] = eV
	fn.Call(in)
	return r
}

func (r *Requester) Bytes() ([]byte, error) {
	_, err := r.response()
	return r.Buffer.Bytes(), err
}

func (r *Requester) String() (string, error) {
	_, err := r.response()
	return r.Buffer.String(), err
}

func (r *Requester) JSON(v interface{}) (err error) {
	_, err = r.response()
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Buffer.Bytes(), v)
	return
}

// Status Code
func (r *Requester) ValidateStatusCode(statusCodes ...int) (b bool) {
	r.response()
	if len(statusCodes) > 0 {
		for _, c := range statusCodes {
			if b = r.StatusCode == c; b {
				break
			}
		}
		// 200 <= x < 300
	} else if r.StatusCode >= 200 && r.StatusCode < 300 {
		b = true
	}

	if !b {
		r.err = errors.New("http: invalid status code " + strconv.Itoa(r.StatusCode))
	}
	return
}

func (r *Requester) Do() (*bytes.Buffer, error) {
	// lazy create request
	r.req = &http.Request{
		Method: r.Method,
		URL:    r.Url,
		Header: r.Header,
	}

	// uploading before
	if r.IsUpload {
		fields, _ := r.rawBody.(map[string]string)
		r.Form(r.files, fields)
	}

	if len(r.cookies) > 0 {
		for _, c := range r.cookies {
			r.req.AddCookie(c)
		}
	}

	// pack body
	r.packBody()

	// uploading after
	if r.IsUpload && r.pg != nil {
		r.pg.Total = r.Length
		r.Body = ioutil.NopCloser(syncreader.New(r.Body, r.pg))
	}

	if r.Body != nil {
		if r.Header.Get("Content-Type") == "" {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		r.req.Body = r.Body
		if r.Length > 0 {
			r.req.ContentLength = r.Length
		}
	}

	var (
		res *http.Response
		err error
	)

	switch r.Url.Scheme {
	case "http", "https":
		r.transport = new(http.Transport)
		if r.tlsconfig != nil {
			r.transport.TLSClientConfig = r.tlsconfig
		}
		r.client = &http.Client{Transport: r.transport, Timeout: r.timeout}
		res, err = r.client.Do(r.req)
		if err != nil {
			return nil, err
		}
		break
	default:
		// TODO @fundon, customzie
		conn, err := net.Dial(r.Url.Scheme, r.Url.Host)
		if err != nil {
			return nil, err
		}
		r.clientConn = httputil.NewClientConn(conn, nil)
		res, err = r.clientConn.Do(r.req)
		if err != nil {
			return nil, err
		}
	}

	defer res.Body.Close()

	r.res = res
	r.StatusCode = res.StatusCode
	// hack, return response header
	r.Header = res.Header
	r.Buffer = new(bytes.Buffer)
	dw := io.MultiWriter(r.Buffer)

	// downloading
	if r.IsDownload {
		var fw io.Writer
		switch t := r.destination.(type) {
		case string:
			p, err := filepath.Abs(t)
			if err != nil {
				return nil, err
			}
			f, err := os.Create(p)
			defer f.Close()
			if err != nil {
				return nil, err
			}
			fw = f
			break
		default:
			fw, _ = t.(io.Writer)
		}
		if r.pg != nil {
			r.pg.Total = res.ContentLength
		}
		dw = io.MultiWriter(dw, r.pg, fw)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.Copy(dw, res.Body)
	if err != nil {
		return nil, err
	}
	return r.Buffer, nil
}

func (r *Requester) TLSConfig(t *tls.Config) *Requester {
	r.tlsconfig = t
	return r
}

func (r *Requester) Cancel() {
	if r.client != nil {
		r.client.Transport.(*http.Transport).CancelRequest(r.req)
	} else if r.clientConn != nil {
		r.clientConn.Close()
	}
}

func (r *Requester) Println() string {
	s := []string{r.Method, r.Url.String()}

	if r.res != nil {
		s = append(s, strconv.Itoa(r.StatusCode))
	}

	return strings.Join(s, " ")
}

func (r *Requester) DebugPrintln() string {
	s := []string{"$ curl -i"}

	if r.Method != GET {
		s = append(s, "-X "+r.Method)
	}

	if r.Url.User != nil {
		s = append(s, "-u "+r.Url.User.String())
	}

	if r.tlsconfig != nil {
		if r.tlsconfig.InsecureSkipVerify {
			s = append(s, "-k")
		} else if r.tlsconfig.RootCAs != nil {
			s = append(s, "--cacert ca.pem")
		}
		s = append(s, "--cert cert.pem")
		s = append(s, "--key key.pem")
	}

	if len(r.cookies) > 0 {
		var cs []string
		for _, c := range r.cookies {
			cs = append(cs, c.String())
		}
		s = append(s, "-b "+strconv.Quote(strings.Join(cs, "; ")))
	}

	for field, _ := range r.Header {
		s = append(s, "-H "+strconv.Quote(field+": "+r.Header.Get(field)))
	}

	if body, _, err := packBody(r.rawBody); err == nil && body != nil {
		b := new(bytes.Buffer)
		io.Copy(b, body)
		s = append(s, "-d "+strconv.Quote(b.String()))
	}

	s = append(s, strconv.Quote(r.Url.String()))

	return strings.Join(s, " \\\n\t")
}

func (r *Requester) Pipe() {
}

func encodesParametersInURL(method string) bool {
	switch method {
	case GET, HEAD, DELETE:
		return true
	}
	return false
}
