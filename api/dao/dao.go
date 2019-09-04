package dao

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"article/api/util"
)

type Dao struct {
	database *sql.DB
}

func (d *Dao) GetArticleDao() *sql.Rows {

	results, err := d.database.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}

	return results
}

func (d *Dao) GetSingleArticleDao(c *gin.Context) (util.Article, *sql.Rows) {
	id := c.Params.ByName("id")
	article := util.Article{}
	errArticle := d.database.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	if errArticle != nil {
		panic(errArticle.Error())
	}
	rows, errImage := d.database.Query("SELECT image_name FROM images WHERE article_uuid  = ?", article.UUID)
	if errImage != nil {
		panic(errImage.Error())
	}

	return article, rows
}

func (d *Dao) DeleteArticleDao(c *gin.Context) {
	id := c.Params.ByName("id")

	article := util.Article{}
	errArticle := d.database.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	if errArticle != nil {
		panic(errArticle.Error())
	}

	var imageNames []util.ImageName

	rows, errImage := d.database.Query("SELECT image_name FROM images WHERE article_uuid  = ?", article.UUID)
	if errImage != nil {
		panic(errImage.Error())
	}

	for rows.Next() {
		imageName := util.ImageName{}
		err := rows.Scan(&imageName.NAME)
		if err != nil {
			panic(err.Error())
		}
		imageNames = append(imageNames, imageName)
	}

	tx, err := d.database.Begin()
	if err != nil {
		fmt.Printf("Failed to begin transaction : %s", err)
		return
	}

	_, deleteArticleErr := tx.Exec("DELETE FROM articles WHERE id= ?", id)
	if deleteArticleErr != nil {
		log.Fatalf("db.Exec(): %s\n", deleteArticleErr)
	}

	_, errDeleteImage := tx.Exec("DELETE FROM images WHERE article_uuid= ?", article.UUID)
	if errDeleteImage != nil {
		log.Fatalf("db.Exec(): %s\n", errDeleteImage)
	}
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return
	}
	tx.Commit()
	DeleteS3Image(imageNames)

}

func (d *Dao) PostDao(article util.Article, uu string) {
	ins, err := d.database.Prepare("INSERT INTO articles(uuid, title,content) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(uu, article.TITLE, article.CONTENT)
}

func (d *Dao) PostImageToDBDao(imageData util.ImageData) {
	ins, err := d.database.Prepare("INSERT INTO images(article_uuid, image_name) VALUES(?,?)")

	for _, imageName := range imageData.IMAGENAMES {

		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(imageData.ARTICLEUUID, imageName.NAME)
	}
}

func NewDao(database *sql.DB) *Dao {
	return &Dao{database: database}
}
