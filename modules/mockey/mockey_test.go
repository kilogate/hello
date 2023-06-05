package mockey

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
)

func TestMockFunc(t *testing.T) {
	mockey.PatchConvey("TestMockFunc", t, func() { // `PatchConvey`外自动释放mock
		mockey.Mock(Foo).Return("c").Build() // mock函数

		res := Foo("a")
		convey.So(res, convey.ShouldEqual, "c")
	})
}

func TestMockMethod(t *testing.T) {
	mockey.PatchConvey("TestMockMethod", t, func() { // `PatchConvey`外自动释放mock
		mockey.Mock(A.Foo).Return("c").Build() // mock方法

		res := new(A).Foo("b")
		convey.So(res, convey.ShouldEqual, "c")
	})
}

func TestMockValue(t *testing.T) {
	mockey.PatchConvey("TestMockXXX", t, func() { // `PatchConvey`外自动释放mock
		mockey.MockValue(&Bar).To(1) // mock变量

		res := Bar
		convey.So(res, convey.ShouldEqual, 1)
	})
}
