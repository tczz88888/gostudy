package main

import (
	"geeweb/day1-http-base/base3/gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"), c.Path)
	})
	engine.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run("127.0.0.1:8000")
}
