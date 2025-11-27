package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello gee")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.String(http.StatusOK, "v1 index")
		})
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello v1")
		})
	}
	r.Run(":9999")
}
