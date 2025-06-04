package app

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func InitializeRoute() *gin.Engine {
	// Setup rate: 20 requests per minute
	rate, _ := limiter.NewRateFromFormatted("20-M")
	store := memory.NewStore()
	rateLimiter20M := ginlimiter.NewMiddleware(limiter.New(store, rate))

	r := gin.Default()
	r.LoadHTMLGlob("tmpl/*")
	r.Use(static.Serve("/assets", static.LocalFile("./public", true)))

	r.GET("/", handleHome)
	r.GET("/s/:id", rateLimiter20M, handleDetail)
	r.POST("/s", rateLimiter20M, handleCreate)
	r.POST("/s/:id/update", rateLimiter20M, handleUpdate)
	r.GET("/api/s/:id", rateLimiter20M, handleDetailAPI)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return r
}
