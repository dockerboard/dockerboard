package quest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	goquery "github.com/google/go-querystring/query"
)

func nopCloser(r io.Reader) io.ReadCloser {
	return ioutil.NopCloser(r)
}

func packBodyByString(s string) (io.ReadCloser, int64) {
	return nopCloser(bytes.NewBufferString(s)), int64(len(s))
}

func packBodyByBytes(b []byte) (io.ReadCloser, int64) {
	return nopCloser(bytes.NewBuffer(b)), int64(len(b))
}

func packBodyByBytesBuffer(b *bytes.Buffer) (io.ReadCloser, int64) {
	return nopCloser(b), int64(b.Len())
}

func packBodyByBytesReader(b *bytes.Reader) (io.ReadCloser, int64) {
	return nopCloser(b), int64(b.Len())
}

func packBodyByPipeReader(pr *io.PipeReader) (io.ReadCloser, int64) {
	b := new(bytes.Buffer)
	length, _ := b.ReadFrom(pr)
	return nopCloser(b), length
}

func packBodyByReader(pr io.Reader) (io.ReadCloser, int64) {
	b := new(bytes.Buffer)
	length, _ := b.ReadFrom(pr)
	return nopCloser(b), length
}

func packBodyByStringsReader(b *strings.Reader) (io.ReadCloser, int64) {
	return nopCloser(b), int64(b.Len())
}

// Pack Request's body to io.ReadCloser.
func packBody(data interface{}) (rc io.ReadCloser, n int64, err error) {
	switch t := data.(type) {
	case nil:
		return
	case string:
		rc, n = packBodyByString(t)
		return
	case []byte:
		rc, n = packBodyByBytes(t)
		return
	case *url.Values:
		rc, n = packBodyByString(t.Encode())
		return
	case *bytes.Buffer:
		rc, n = packBodyByBytesBuffer(t)
		return
	case *bytes.Reader:
		rc, n = packBodyByBytesReader(t)
		return
	case *strings.Reader:
		rc, n = packBodyByStringsReader(t)
		return
	case *io.PipeReader:
		rc, n = packBodyByPipeReader(t)
		return
	// JSON Object
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, 0, err
		}
		rc, n = packBodyByBytes(b)
	}
	return
}

func QueryString(options interface{}) (qs string, err error) {
	switch t := options.(type) {
	case string:
		qs = url.QueryEscape(t)
		return
	case []byte:
		qs = url.QueryEscape(string(t))
		return
	case *url.Values:
		qs = t.Encode()
		return
	default:
		v, err := goquery.Values(t)
		if err != nil {
			return "", err
		}
		qs = v.Encode()
	}
	return
}

func getFile(f interface{}) (file *os.File, name string, err error) {
	switch t := f.(type) {
	case string:
		name = t
		file, err = os.Open(name)
		break
	case *os.File:
		var fs os.FileInfo
		file = t
		fs, err = t.Stat()
		name = fs.Name()
		break
	case os.FileInfo:
		name = t.Name()
		file, err = os.Open(name)
		break
	default:
		err = errors.New("Not a file.")
	}
	if err != nil {
		file = nil
		name = ""
	}
	return
}
