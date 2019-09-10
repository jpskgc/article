package main

import (
	"article/api/dao"
	"article/api/db"
	"article/api/s3"
	"article/api/service"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"

	"article/api/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	}

	db := db.NewDatabase(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	s3 := s3.NewS3(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	dao := dao.NewDao(db.DATABASE, s3)
	service := service.NewService(dao)
	cntlr := controller.NewController(service)

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

	router.Run(":2345")
}
