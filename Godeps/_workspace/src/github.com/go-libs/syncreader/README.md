
# syncreader

SyncReader creates a reader that its reads to synchronize with first reader, similar to the `io.MultiWriter `.

View the [docs][].

## APIs

### `New(readers ...io.Reader)`

`First reader` is `source reader`, other readers read to synchronize whith first reader.
If `first readera` retures `EOF`, other readers also return.


## Usage

Reading a file in progress.

```go
import "github.com/go-libs/syncreader"
import "github.com/go-libs/progress"

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
r := syncreader.New(f, p)
_, err = b.ReadFrom(r)
if err != nil {
  log.Fatalln(err)
}
```


View the [docs][].


[docs]: http://godoc.org/github.com/go-libs/syncreader
