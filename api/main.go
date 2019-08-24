package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"

	"github.com/jpskgc/article/api/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	}
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/article")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS article;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use article;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `articles` (`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,uuid varchar(36), `title` VARCHAR(100) NOT NULL,`content` TEXT NOT NULL ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table IF NOT EXISTS images (id int AUTO_INCREMENT NOT NULL PRIMARY KEY, article_uuid varchar(36), image_name varchar(50));	")
	if err != nil {
		panic(err)
	}

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
			controller.GetArticleController(c, db)
		})
		api.GET("/article/:id", func(c *gin.Context) {
			controller.GetSingleArticleController(c, db)
		})
		api.GET("/delete/:id", func(c *gin.Context) {
			controller.DeleteArticleController(c, db)
		})
		api.POST("/post", func(c *gin.Context) {
			controller.PostController(c, db)
		})
		api.POST("/post/image", func(c *gin.Context) {
			controller.PostImageController(c)
		})
		api.POST("/post/image/db", func(c *gin.Context) {
			controller.PostImageToDBController(c, db)
		})
	}

	router.Run(":2345")
}
