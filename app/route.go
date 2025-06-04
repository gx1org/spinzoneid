package app

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitializeRoute() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("tmpl/*")
	r.Use(static.Serve("/assets", static.LocalFile("./public", true)))

	r.GET("/", handleHome)
	r.GET("/s/:id", handleDetail)

	r.POST("/s", handleCreate)
	r.POST("/s/:id/update", handleUpdate)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return r
}
