package gomonkey

import (
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyFuncVar
func TestApplyFuncVar(t *testing.T) {
	convey.Convey("TestApplyFuncVar", t, func() {

		convey.Convey("for success", func() {
			str := "hello"
			patches := gomonkey.ApplyFuncVar(&fake.Marshal, func(_ interface{}) ([]byte, error) {
				return []byte(str), nil
			})
			defer patches.Reset()

			bytes, err := fake.Marshal(nil)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(string(bytes), convey.ShouldEqual, str)
		})

		convey.Convey("for fail", func() {
			patches := gomonkey.ApplyFuncVar(&fake.Marshal, func(_ interface{}) ([]byte, error) {
				return nil, fake.ErrActual
			})
			defer patches.Reset()

			_, err := fake.Marshal(nil)
			convey.So(err, convey.ShouldEqual, fake.ErrActual)
		})

	})
}
