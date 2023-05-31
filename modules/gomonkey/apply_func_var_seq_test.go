package gomonkey

import (
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyFuncVarSeq
func TestApplyFuncVarSeq(t *testing.T) {
	convey.Convey("TestApplyFuncVarSeq", t, func() {

		convey.Convey("default times is 1", func() {
			info1 := "hello cpp"
			info2 := "hello golang"
			info3 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{[]byte(info1), nil}},
				{Values: gomonkey.Params{[]byte(info2), nil}},
				{Values: gomonkey.Params{[]byte(info3), nil}},
			}
			patches := gomonkey.ApplyFuncVarSeq(&fake.Marshal, outputs)
			defer patches.Reset()

			bytes, err := fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info1)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info2)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info3)
		})

		convey.Convey("retry succ util the third times", func() {
			info1 := "hello cpp"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{[]byte(""), fake.ErrActual}, Times: 2},
				{Values: gomonkey.Params{[]byte(info1), nil}},
			}
			patches := gomonkey.ApplyFuncVarSeq(&fake.Marshal, outputs)
			defer patches.Reset()

			bytes, err := fake.Marshal("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info1)
		})

		convey.Convey("batch operations failed on the third time", func() {
			info1 := "hello gomonkey"
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{[]byte(info1), nil}, Times: 2},
				{Values: gomonkey.Params{[]byte(""), fake.ErrActual}},
			}
			patches := gomonkey.ApplyFuncVarSeq(&fake.Marshal, outputs)
			defer patches.Reset()

			bytes, err := fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info1)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, info1)

			bytes, err = fake.Marshal("")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
		})

	})
}
