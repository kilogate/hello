package main

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"

	"code.byted.org/middleware/hertz/byted"
)

// go run modules/hertz/hertz.go -psm xxx -conf-dir modules/hertz -log-dir ./log
func main() {
	byted.Init()
	r := byted.Default()

	r.GET("/test/get", handleGet)
	r.PUT("/test/put", handlePut)
	r.POST("/test/post", handlePost)
	r.DELETE("/test/delete/:id", handleDelete)

	r.Spin()
}

func handleGet(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(200, struct {
		ID   string
		Name string
	}{
		ID:   id,
		Name: name,
	})
}

func handlePut(ctx context.Context, c *app.RequestContext) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	c.JSON(200, struct {
		ID   string
		Name string
	}{
		ID:   id,
		Name: name,
	})
}

func handlePost(ctx context.Context, c *app.RequestContext) {
	var s = struct {
		ID   string
		Name string
	}{}
	_ = json.Unmarshal(c.Request.Body(), &s)
	c.JSON(200, s)
}

func handleDelete(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	c.JSON(200, struct {
		ID   string
		Name string
	}{
		ID:   id,
		Name: "",
	})
}
