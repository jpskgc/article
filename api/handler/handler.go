package handler

import (
	"article/api/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router     *gin.Engine
	controller controller.ControllerInterface
}

func NewHandler(cntlr controller.ControllerInterface) *Handler {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	api := router.Group("/api")
	{
		api.GET("/articles", func(c *gin.Context) {
			cntlr.GetArticleController(c)
		})
		api.GET("/article/:id", func(c *gin.Context) {
			cntlr.GetSingleArticleController(c)
		})
		api.GET("/delete/:id", func(c *gin.Context) {
			cntlr.DeleteArticleController(c)
		})
		api.POST("/post", func(c *gin.Context) {
			cntlr.PostController(c)
		})
		api.POST("/post/image", func(c *gin.Context) {
			cntlr.PostImageController(c)
		})
		api.POST("/post/image/db", func(c *gin.Context) {
			cntlr.PostImageToDBController(c)
		})
	}

	Handler := new(Handler)
	Handler.Router = router
	return Handler
}
