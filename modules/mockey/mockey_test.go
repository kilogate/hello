package mockey

import (
	"fmt"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
)

func TestMockFunc(t *testing.T) {
	convey.Convey("TestMockFunc", t, func() {
		mockey.Mock(Foo).Return("c").Build()

		res := Foo("a")
		convey.So(res, convey.ShouldEqual, "c")
	})
}

func TestMockMethod(t *testing.T) {
	convey.Convey("TestMockMethod", t, func() {
		mockey.Mock((*A).Foo).Return("c").Build()

		res := new(A).Foo("b")
		convey.So(res, convey.ShouldEqual, "c")
	})
}

func TestMockValue(t *testing.T) {
	convey.Convey("TestMockXXX", t, func() {
		mockey.MockValue(&Bar).To(1)

		res := Bar
		convey.So(res, convey.ShouldEqual, 1)
	})
}

func TestPatchConvey(t *testing.T) {
	mockey.PatchConvey("TestPatchConvey", t, func() {
		mockey.Mock(Foo).Return("c").Build()

		res := Foo("a")
		convey.So(res, convey.ShouldEqual, "c")
	})

	// `PatchConvey`外自动释放mock
	res := Foo("a")
	fmt.Println(res) // a
}
