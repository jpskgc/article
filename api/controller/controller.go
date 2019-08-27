package controller

import (
	"database/sql"
	"net/http"

	"article/api/service"

	"github.com/gin-gonic/gin"
)

func GetArticleController(c *gin.Context, db *sql.DB) {
	articles := service.GetArticleService(c, db)
	c.JSON(http.StatusOK, articles)
}

func GetSingleArticleController(c *gin.Context, db *sql.DB) {
	article := service.GetSingleArticleService(c, db)

	c.JSON(http.StatusOK, article)
}

func DeleteArticleController(c *gin.Context, db *sql.DB) {
	service.DeleteArticleService(c, db)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostController(c *gin.Context, db *sql.DB) {
	uu := service.PostService(c, db)

	c.JSON(http.StatusOK, gin.H{"uuid": uu})
}

func PostImageController(c *gin.Context) {

	imageNames := service.PostImageService(c)

	c.JSON(http.StatusOK, imageNames)
}

func PostImageToDBController(c *gin.Context, db *sql.DB) {
	service.PostImageToDBService(c, db)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
