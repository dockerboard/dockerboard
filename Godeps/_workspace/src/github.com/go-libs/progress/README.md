
# progress

Making a progress for `Reader` or `Writer` write/read bytes.

View the [docs][].


## Usage

```go
import "github.com/go-libs/progress"
```


### Writing bytes in progress

```go
import "github.com/go-libs/progress"

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

p := Progress.New()
p.Total = fs.Size()
p.Progress = func(current, total, expected int64) {
  log.Println("Writing", current, total, expected)
}

b := new(bytes.Buffer)
w := io.MultiWriter(p, b)
_, err = io.Copy(w, f)
if err != nil {
  log.Fatalln(err)
}
```


### Reading bytes in progress


```go
import "github.com/go-libs/syncreader"
import "github.com/go-libs/progress"

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

p := Progress.New()
p.Total = fs.Size()
p.Progress = func(current, total, expected int64) {
  log.Println("Reading", current, total, expected)
}

b := new(bytes.Buffer)
r := syncreader.New(f, p)
_, err = b.ReadFrom(r)
if err != nil {
  log.Fatalln(err)
}
```


### Using `IgnoreTotal` flag to ignore total size in `Progress` initialization.

```go
p := Progress.New()
p.IgnoreTotal = true
p.Progress = func(current, total, expected int64) {
  log.Println("Reading", current, total, expected)
}
```


[docs]: http://godoc.org/github.com/go-libs/progress

