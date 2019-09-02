package service

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"article/api/dao"

	"article/api/util"
)

func GetArticleService(c *gin.Context, db *sql.DB) []util.Article {
	var articles []util.Article

	results := dao.GetArticleDao(db)

	article := util.Article{}
	for results.Next() {
		err := results.Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
		if err != nil {
			panic(err.Error())
		} else {
			articles = append(articles, article)
		}
	}
	return articles
}

func GetSingleArticleService(c *gin.Context, db *sql.DB) util.Article {

	article, rows := dao.GetSingleArticleDao(c, db)

	for rows.Next() {
		imageName := util.ImageName{}
		err := rows.Scan(&imageName.NAME)
		if err != nil {
			panic(err.Error())
		}
		article.IMAGENAMES = append(article.IMAGENAMES, imageName)
	}

	return article
}

func DeleteArticleService(c *gin.Context, db *sql.DB) {
	dao.DeleteArticleDao(c, db)
}

func PostService(c *gin.Context, db *sql.DB) string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	uu := u.String()
	var article util.Article
	c.BindJSON(&article)
	dao.PostDao(db, article, uu)

	return uu
}

func PostImageService(c *gin.Context) []util.ImageName {
	return dao.PostImageToS3(c)
}

func PostImageToDBService(c *gin.Context, db *sql.DB) {
	var imageData util.ImageData
	c.BindJSON(&imageData)
	dao.PostImageToDBDao(imageData, db)
}
