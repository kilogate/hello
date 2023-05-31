package gomonkey

import (
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
)

var num = 10

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestApplyGlobalVar
func TestApplyGlobalVar(t *testing.T) {
	convey.Convey("TestApplyGlobalVar", t, func() {

		convey.Convey("change", func() {
			patches := gomonkey.ApplyGlobalVar(&num, 150)
			defer patches.Reset()
			convey.So(num, convey.ShouldEqual, 150)
		})

		convey.Convey("recover", func() {
			convey.So(num, convey.ShouldEqual, 10)
		})

	})
}
