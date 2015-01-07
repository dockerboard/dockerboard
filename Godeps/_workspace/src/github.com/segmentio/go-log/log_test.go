package log

// import "github.com/bmizerany/assert"
import "testing"
import "bytes"
import "fmt"
import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: real tests :)
func TestLog(t *testing.T) {
	l := New(os.Stderr, DEBUG, "")
	l.Debug("something happened")
	l.Info("hello %s", "Tobi")
	l.Error("boom something exploded")

	l.SetPrefix("myapp")
	l.Info("something")
	l.Info("something else")
	l.Info("moar stuff here")

	Debug("something")
	Emergency("hello %s %s", "tobi", "ferret")

	bytes.NewBufferString("foo\nbar %s").WriteTo(Log)

	Check(fmt.Errorf("something exploded"))
}
