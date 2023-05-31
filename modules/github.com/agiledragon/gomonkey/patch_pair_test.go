package gomonkey

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
	"github.com/smartystreets/goconvey/convey"
)

// cd /Users/bytedance/IdeaProjects/hello/modules/github.com/agiledragon/gomonkey && GOARCH=amd64 go test -v -run=TestPatchPair
func TestPatchPair(t *testing.T) {
	convey.Convey("TestPatchPair", t, func() {
		patchPairs := [][2]interface{}{
			{
				fake.Exec,
				func(_ string, _ ...string) (string, error) {
					return expectOutput, nil
				},
			},
			{
				json.Unmarshal,
				func(_ []byte, v interface{}) error {
					p := v.(*map[int]int)
					*p = make(map[int]int)
					(*p)[1] = 2
					(*p)[2] = 4
					return nil
				},
			},
		}
		patches := gomonkey.NewPatches()
		defer patches.Reset()
		for _, pair := range patchPairs {
			patches.ApplyFunc(pair[0], pair[1])
		}

		output, err := fake.Exec("", "")
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(output, convey.ShouldEqual, expectOutput)

		var m map[int]int
		err = json.Unmarshal(nil, &m)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(m[1], convey.ShouldEqual, 2)
		convey.So(m[2], convey.ShouldEqual, 4)
	})
}
