package gomonkey

import (
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyFuncSeq -gcflags=-l
func TestApplyFuncSeq(t *testing.T) {
	convey.Convey("TestApplyFuncSeq", t, func() {

		convey.Convey("default times is 1", func() {
			info1 := "hello cpp"
			info2 := "hello golang"
			info3 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{info1, nil}},
				{Values: gomonkey.Params{info2, nil}},
				{Values: gomonkey.Params{info3, nil}},
			}
			patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
			defer patches.Reset()

			output, err := fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info2)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info3)
		})

		convey.Convey("retry success util the third times", func() {
			info1 := "hello cpp"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{"", fake.ErrActual}, Times: 2},
				{Values: gomonkey.Params{info1, nil}},
			}
			patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
			defer patches.Reset()

			output, err := fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)
		})

		convey.Convey("batch operations failed on the third time", func() {
			info1 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{info1, nil}, Times: 2},
				{Values: gomonkey.Params{"", fake.ErrActual}},
			}
			patches := gomonkey.ApplyFuncSeq(fake.ReadLeaf, outputs)
			defer patches.Reset()

			output, err := fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, info1)

			output, err = fake.ReadLeaf("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
		})

	})
}
