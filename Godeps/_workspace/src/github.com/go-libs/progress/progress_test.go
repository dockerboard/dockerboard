package progress

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/go-libs/syncreader"
)

func TestNew(t *testing.T) {
	p := New()
	zero := int64(0)
	assert.Equal(t, zero, p.Current)
	assert.Equal(t, zero, p.Total)
	assert.Equal(t, zero, p.Expected)
}

func TestProgressWriter(t *testing.T) {
	filename := "progress_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := New()
	p.Total = fs.Size()
	p.Progress = func(current, total, expected int64) {
		log.Println("Writing", current, total, expected)
		assert.Equal(t, true, current <= total)
	}

	b := new(bytes.Buffer)
	w := io.MultiWriter(p, b)
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, fs.Size(), int64(b.Len()))
}

func TestProgressWriterIgnoreTotal(t *testing.T) {
	filename := "progress_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := New()
	p.IgnoreTotal = true
	p.Progress = func(current, total, expected int64) {
		log.Println("Ignore total writing", current, total, expected)
		assert.Equal(t, true, current >= total)
	}

	b := new(bytes.Buffer)
	w := io.MultiWriter(p, b)
	_, err = io.Copy(w, f)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, fs.Size(), int64(b.Len()))
}

func TestProgressReader(t *testing.T) {
	filename := "progress_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := New()
	p.Total = fs.Size()
	p.Progress = func(current, total, expected int64) {
		log.Println("Reading", current, total, expected)
		assert.Equal(t, true, current <= total)
	}

	b := new(bytes.Buffer)
	r := syncreader.New(f, p)
	_, err = b.ReadFrom(r)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, fs.Size(), int64(b.Len()))
}

func TestProgressReaderIgnoreTotal(t *testing.T) {
	filename := "progress_test.go"
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fs, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p := New()
	p.IgnoreTotal = true
	p.Progress = func(current, total, expected int64) {
		log.Println("Ignore total reading", current, total, expected)
		assert.Equal(t, true, current >= total)
	}

	b := new(bytes.Buffer)
	r := syncreader.New(f, p)
	_, err = b.ReadFrom(r)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, fs.Size(), int64(b.Len()))
}
