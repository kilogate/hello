package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyMethodSeq -gcflags=-l
func TestApplyMethodSeq(t *testing.T) {
	e := &fake.Etcd{}

	convey.Convey("TestApplyMethodSeq", t, func() {

		convey.Convey("default times is 1", func() {
			info1 := "hello cpp"
			info2 := "hello golang"
			info3 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{info1, nil}},
				{Values: gomonkey.Params{info2, nil}},
				{Values: gomonkey.Params{info3, nil}},
			}
			patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(e), "Retrieve", outputs)
			defer patches.Reset()

			output, err := e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info2)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info3)
		})

		convey.Convey("retry success util the third times", func() {
			info1 := "hello cpp"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{"", fake.ErrActual}, Times: 2},
				{Values: gomonkey.Params{info1, nil}},
			}
			patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(e), "Retrieve", outputs)
			defer patches.Reset()

			output, err := e.Retrieve("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)
		})

		convey.Convey("batch operations failed on the third time", func() {
			info1 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{info1, nil}, Times: 2},
				{Values: gomonkey.Params{"", fake.ErrActual}},
			}
			patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(e), "Retrieve", outputs)
			defer patches.Reset()

			output, err := e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = e.Retrieve("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
		})

	})
}
