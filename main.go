package main

import (
	"fmt"
	"gee/gee"
	"log"
	"net/http"
	"time"
)

// 一个中间件logger
func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		//start time
		t := time.Now()
		//将控制权交给下一个中间件或最终的处理函数
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
func OnlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Println("only for v2 正在检查权限...,Auth passed")
		c.Next()
	}
}
func main() {
	r := gee.New()
	r.Use(Logger())
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
	v2 := r.Group("/v2")
	v2.Use(OnlyForV2())
	{
		v2.GET("/", func(c *gee.Context) {
			c.String(http.StatusOK, "v2 index")
		})
		v2.GET("/hello", func(c *gee.Context) {
			fmt.Println("v2 hello")
			c.String(http.StatusOK, "hello v2")
		})
	}
	r.Run(":9999")
}
