package gomonkey

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyFunc
func TestApplyFunc(t *testing.T) {
	convey.Convey("TestApplyFunc", t, func() {

		convey.Convey("one func for success", func() {
			patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return expectOutput, nil
			})
			defer patches.Reset()

			output, err := fake.Exec("", "")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, expectOutput)
		})

		convey.Convey("one func for fail", func() {
			patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return "", fake.ErrActual
			})
			defer patches.Reset()

			output, err := fake.Exec("", "")
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
			convey.So(output, convey.ShouldEqual, "")
		})

		convey.Convey("two functions", func() {
			patches := gomonkey.ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return expectOutput, nil
			})
			defer patches.Reset()
			patches.ApplyFunc(fake.Belong, func(_ string, _ []string) bool {
				return true
			})

			output, err := fake.Exec("", "")
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(output, convey.ShouldEqual, expectOutput)

			flag := fake.Belong("", nil)
			convey.So(flag, convey.ShouldBeTrue)
		})

		convey.Convey("input and output param", func() {
			patches := gomonkey.ApplyFunc(json.Unmarshal, func(_ []byte, v interface{}) error {
				p := v.(*map[int]int)
				*p = make(map[int]int)
				(*p)[1] = 2
				(*p)[2] = 4
				return nil
			})
			defer patches.Reset()

			var m map[int]int
			err := json.Unmarshal(nil, &m)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(m[1], convey.ShouldEqual, 2)
			convey.So(m[2], convey.ShouldEqual, 4)
		})

	})
}
