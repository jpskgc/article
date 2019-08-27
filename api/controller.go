package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticleController(c *gin.Context, db *sql.DB) {
	articles := GetArticleService(c, db)
	c.JSON(http.StatusOK, articles)
}

func GetSingleArticleController(c *gin.Context, db *sql.DB) {
	article := GetSingleArticleService(c, db)

	c.JSON(http.StatusOK, article)
}

func DeleteArticleController(c *gin.Context, db *sql.DB) {
	DeleteArticleService(c, db)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostController(c *gin.Context, db *sql.DB) {
	uu := PostService(c, db)

	c.JSON(http.StatusOK, gin.H{"uuid": uu})
}

func PostImageController(c *gin.Context) {

	imageNames := PostImageService(c)

	c.JSON(http.StatusOK, imageNames)
}

func PostImageToDBController(c *gin.Context, db *sql.DB) {
	PostImageToDBService(c, db)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
