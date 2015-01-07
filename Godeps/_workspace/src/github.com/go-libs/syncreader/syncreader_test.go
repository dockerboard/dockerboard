package syncreader

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/go-libs/progress"
)

func TestNew(t *testing.T) {
	filename := "syncreader_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := progress.New()
	p.Total = fs.Size()
	p.Progress = func(c, t, e int64) {
		log.Println(c, t, e)
	}
	b := new(bytes.Buffer)
	r := New(f, p)
	_, err = b.ReadFrom(r)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, fs.Size(), int64(b.Len()))
}
