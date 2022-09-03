package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type routes struct {
	router *gin.Engine
}

func NewRouters() routes {
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"POST", "GET"},
			MaxAge:       12 * time.Hour,
		}))

	api := r.router.Group("/api")
	api.GET("/getPosts", GetPosts)
	api.GET("/getPost", GetPost)
	api.POST("/writePost", WritePost)
	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
