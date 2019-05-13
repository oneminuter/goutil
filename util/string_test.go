package ut

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStr2bytes(t *testing.T) {
	Convey("Str2bytes", t, func() {
		So(string(Str2bytes("abc")), ShouldEqual, "abc")
	})
}
