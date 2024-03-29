package dao

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	"article/api/s3"
	"article/api/util"
)

type Dao struct {
	database *sql.DB
	s3       s3.S3Interface
}

func NewDao(database *sql.DB, s3 s3.S3Interface) *Dao {
	objs := &Dao{database: database, s3: s3}
	return objs
}

type DaoInterface interface {
	GetArticleDao() *sql.Rows
	GetSingleArticleDao(id string) (util.Article, *sql.Rows)
	DeleteArticleDao(id string)
	PostDao(article util.Article, uu string)
	PostImageToDBDao(imageData util.ImageData)
	PostImageToS3Dao(file *multipart.FileHeader, imageName string)
}

func (d *Dao) GetArticleDao() *sql.Rows {

	results, err := d.database.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}

	return results
}

func (d *Dao) GetSingleArticleDao(id string) (util.Article, *sql.Rows) {
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

func (d *Dao) DeleteArticleDao(id string) {

	article := util.Article{}
	errArticle := d.database.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	if errArticle != nil {
		panic(errArticle.Error())
	}

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
		d.s3.DeleteS3Image(imageName)
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

func (d *Dao) PostImageToS3Dao(file *multipart.FileHeader, imageName string) {
	d.s3.PostImageToS3(file, imageName)
}
