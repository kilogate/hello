package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyMethod
func TestApplyMethod(t *testing.T) {
	slice := fake.NewSlice()
	var s *fake.Slice

	convey.Convey("TestApplyMethod", t, func() {

		convey.Convey("for success", func() {
			err := slice.Add(1)
			convey.So(err, convey.ShouldEqual, nil)

			patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
				return nil
			})
			defer patches.Reset()

			err = slice.Add(1)
			convey.So(err, convey.ShouldEqual, nil)

			err = slice.Remove(1)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(len(slice), convey.ShouldEqual, 0)
		})

		convey.Convey("for already exist", func() {
			err := slice.Add(2)
			convey.So(err, convey.ShouldEqual, nil)

			patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemExsit
			})
			defer patches.Reset()

			err = slice.Add(1)
			convey.So(err, convey.ShouldEqual, fake.ErrElemExsit)

			err = slice.Remove(2)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(len(slice), convey.ShouldEqual, 0)
		})

		convey.Convey("two methods", func() {
			err := slice.Add(3)
			convey.So(err, convey.ShouldEqual, nil)
			defer slice.Remove(3)

			patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemExsit
			})
			defer patches.Reset()
			patches.ApplyMethod(reflect.TypeOf(s), "Remove", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemExsit
			})

			err = slice.Add(2)
			convey.So(err, convey.ShouldEqual, fake.ErrElemExsit)

			err = slice.Remove(1)
			convey.So(err, convey.ShouldEqual, fake.ErrElemExsit)
			convey.So(len(slice), convey.ShouldEqual, 1)
			convey.So(slice[0], convey.ShouldEqual, 3)
		})

		convey.Convey("one func and one method", func() {
			err := slice.Add(4)
			convey.So(err, convey.ShouldEqual, nil)
			defer slice.Remove(4)

			patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return expectOutput, nil
			})
			defer patches.Reset()
			patches.ApplyMethod(reflect.TypeOf(s), "Remove", func(_ *fake.Slice, _ int) error {
				return fake.ErrElemNotExsit
			})

			output, err := fake.Exec("", "")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, expectOutput)

			err = slice.Remove(1)
			convey.So(err, convey.ShouldEqual, fake.ErrElemNotExsit)
			convey.So(len(slice), convey.ShouldEqual, 1)
			convey.So(slice[0], convey.ShouldEqual, 4)
		})

	})
}
