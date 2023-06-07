package main

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/smartystreets/goconvey/convey"

	"code.byted.org/middleware/hertz/pkg/common/config"
	"code.byted.org/middleware/hertz/pkg/route"
	"code.byted.org/middleware/hertz_ext/v2/hertztest"
)

func Test_handleGet(t *testing.T) {
	convey.Convey("Test_handleGet", t, func() {
		r := route.NewEngine(config.NewOptions([]config.Option{}))
		r.GET("/test/get", handleGet)

		w := hertztest.PerformRequest(r, consts.MethodGet, "/test/get?id=123&name=Tom", nil)
		resp := w.Result()

		var s = struct {
			ID   string
			Name string
		}{}
		err := json.Unmarshal(resp.Body(), &s)

		want := struct {
			ID   string
			Name string
		}{
			ID:   "123",
			Name: "Tom",
		}

		convey.So(resp.StatusCode(), convey.ShouldEqual, 200)
		convey.So(string(resp.Header.ContentType()), convey.ShouldEqual, "application/json; charset=utf-8")
		convey.So(err, convey.ShouldBeNil)
		convey.So(s, convey.ShouldResemble, want)
	})
}

func Test_handlePut(t *testing.T) {
	convey.Convey("Test_handlePut", t, func() {
		r := route.NewEngine(config.NewOptions([]config.Option{}))
		r.PUT("/test/put", handlePut)

		values := url.Values{
			"id":   {"123"},
			"name": {"Tom"},
		}
		body := values.Encode()
		w := hertztest.PerformRequest(r, consts.MethodPut, "/test/put", &hertztest.Body{Body: bytes.NewBufferString(body), Len: len(body)},
			hertztest.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"})
		resp := w.Result()

		var s = struct {
			ID   string
			Name string
		}{}
		err := json.Unmarshal(resp.Body(), &s)

		want := struct {
			ID   string
			Name string
		}{
			ID:   "123",
			Name: "Tom",
		}

		convey.So(resp.StatusCode(), convey.ShouldEqual, 200)
		convey.So(string(resp.Header.ContentType()), convey.ShouldEqual, "application/json; charset=utf-8")
		convey.So(err, convey.ShouldBeNil)
		convey.So(s, convey.ShouldResemble, want)
	})
}

func Test_handlePost(t *testing.T) {
	convey.Convey("Test_handlePost", t, func() {
		r := route.NewEngine(config.NewOptions([]config.Option{}))
		r.POST("/test/post", handlePost)

		s := struct {
			ID   string
			Name string
		}{
			ID:   "123",
			Name: "Tom",
		}
		bodyBytes, _ := json.Marshal(s)
		w := hertztest.PerformRequest(r, consts.MethodPost, "/test/post", &hertztest.Body{Body: bytes.NewBuffer(bodyBytes), Len: len(bodyBytes)})
		resp := w.Result()

		err := json.Unmarshal(resp.Body(), &s)

		want := struct {
			ID   string
			Name string
		}{
			ID:   "123",
			Name: "Tom",
		}

		convey.So(resp.StatusCode(), convey.ShouldEqual, 200)
		convey.So(err, convey.ShouldBeNil)
		convey.So(s, convey.ShouldResemble, want)
		convey.So(string(resp.Header.ContentType()), convey.ShouldEqual, "application/json; charset=utf-8")
	})
}

func Test_handleDelete(t *testing.T) {
	convey.Convey("Test_handleDelete", t, func() {
		r := route.NewEngine(config.NewOptions([]config.Option{}))
		r.DELETE("/test/delete/:id", handleDelete)

		w := hertztest.PerformRequest(r, consts.MethodDelete, "/test/delete/123", nil, hertztest.Header{Key: "user", Value: "Tom"})
		resp := w.Result()

		var s = struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Id   string `json:"id"`
				User string `json:"user"`
			} `json:"data"`
		}{}
		err := json.Unmarshal(resp.Body(), &s)

		want := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Id   string `json:"id"`
				User string `json:"user"`
			} `json:"data"`
		}{
			Code:    0,
			Message: "success",
			Data: struct {
				Id   string `json:"id"`
				User string `json:"user"`
			}{
				Id:   "123",
				User: "Tom",
			},
		}

		convey.So(resp.StatusCode(), convey.ShouldEqual, 200)
		convey.So(string(resp.Header.Get("Update-At")), convey.ShouldNotBeEmpty)
		convey.So(err, convey.ShouldBeNil)
		convey.So(s, convey.ShouldResemble, want)
	})
}
