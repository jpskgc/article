package dao

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jpskgc/article/api/util"
)

// type Article struct {
// 	ID         int         `json:"id"`
// 	UUID       string      `json:"uuid"`
// 	TITLE      string      `json:"title"`
// 	CONTENT    string      `json:"content"`
// 	IMAGENAMES []ImageName `json:"imageNames"`
// }

// type ImageName struct {
// 	NAME string `json:"name"`
// }

func GetArticleDao(db *sql.DB) *sql.Rows {
	results, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}

	return results
}

func GetSingleArticleDao(c *gin.Context, db *sql.DB) (util.Article, *sql.Rows) {
	id := c.Params.ByName("id")
	article := util.Article{}
	errArticle := db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	if errArticle != nil {
		panic(errArticle.Error())
	}
	rows, errImage := db.Query("SELECT image_name FROM images WHERE article_uuid  = ?", article.UUID)
	if errImage != nil {
		panic(errImage.Error())
	}

	return article, rows
}

func DeleteArticleDao(c *gin.Context, db *sql.DB) {
	id := c.Params.ByName("id")
	_, err := db.Exec("DELETE FROM articles WHERE id= ?", id)
	if err != nil {
		log.Fatalf("db.Exec(): %s\n", err)
	}
}

func PostDao(db *sql.DB, article util.Article, uu string) {
	ins, err := db.Prepare("INSERT INTO articles(uuid, title,content) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(uu, article.TITLE, article.CONTENT)
}

func PostImageToDBDao(c *gin.Context, db *sql.DB) {
	var imageData util.ImageData
	c.BindJSON(&imageData)

	for _, imageName := range imageData.IMAGENAMES {
		ins, err := db.Prepare("INSERT INTO images(article_uuid, image_name) VALUES(?,?)")
		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(imageData.ARTICLEUUID, imageName.NAME)
	}
}
