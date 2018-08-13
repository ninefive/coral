package routers

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/middleware"
	"github.com/ninefive/coral/routers/api/health"
	"github.com/ninefive/coral/routers/api/v1/user"
)

func Init(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//swagger api docs
	//TODO

	//pprof
	pprof.Register(g)

	g.POST("/login", user.Login)

	u := g.Group("/u/v1/user")
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	h := g.Group("/health")
	{
		h.GET("/health", health.HealthCheck)
		h.GET("/disk", health.DiskCheck)
		h.GET("/cpu", health.CPUCheck)
		h.GET("/mem", health.RAMCheck)
	}

	return g
}
