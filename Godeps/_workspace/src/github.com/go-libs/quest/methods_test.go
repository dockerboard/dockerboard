package quest

import (
	"testing"

	mocha "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	mocha.Convey("Methods", t, func() {
		mocha.So("OPTIONS", mocha.ShouldEqual, OPTIONS)
		mocha.So("PUT", mocha.ShouldEqual, PUT)
	})
}
