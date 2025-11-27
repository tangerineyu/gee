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
	r.GET("/json", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"title": "Day 2",
			"frame": "gee",
		})
	})
	r.GET("/html", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Run(":9999")
}
