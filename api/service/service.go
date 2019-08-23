package service

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/jpskgc/article/api/dao"

	"github.com/jpskgc/article/api/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

// type ImageData struct {
// 	ARTICLEUUID string      `json:"articleUUID"`
// 	IMAGENAMES  []ImageName `json:"imageNames"`
// }

// type Param struct {
// 	Bucket      string
// 	Key         string
// 	Expires     string
// 	ContentType string
// }

func GetArticleService(c *gin.Context, db *sql.DB) []util.Article {
	var articles []util.Article

	results := dao.GetArticleDao(db)
	// results, err := db.Query("SELECT * FROM articles")
	// if err != nil {
	// 	panic(err.Error())
	// }
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
	// id := c.Params.ByName("id")
	// article := Article{}
	// errArticle := db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	// if errArticle != nil {
	// 	panic(errArticle.Error())
	// }
	// rows, errImage := db.Query("SELECT image_name FROM images WHERE article_uuid  = ?", article.UUID)
	// if errImage != nil {
	// 	panic(errImage.Error())
	// }
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
	// id := c.Params.ByName("id")
	// _, err := db.Exec("DELETE FROM articles WHERE id= ?", id)
	// if err != nil {
	// 	log.Fatalf("db.Exec(): %s\n", err)
	// }
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
	// ins, err := db.Prepare("INSERT INTO articles(uuid, title,content) VALUES(?,?,?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//ins.Exec(uu, article.TITLE, article.CONTENT)

	return uu
}

func PostImageService(c *gin.Context) []util.ImageName {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	form, _ := c.MultipartForm()

	files := form.File["images[]"]

	var imageNames []util.ImageName
	imageName := util.ImageName{}

	for _, file := range files {

		f, err := file.Open()

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		size := file.Size
		buffer := make([]byte, size)

		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()

		f.Read(buffer)
		fileBytes := bytes.NewReader(buffer)
		fileType := http.DetectContentType(buffer)
		path := "/media/" + uu
		params := &s3.PutObjectInput{
			Bucket:        aws.String("article-s3-jpskgc"),
			Key:           aws.String(path),
			Body:          fileBytes,
			ContentLength: aws.Int64(size),
			ContentType:   aws.String(fileType),
		}
		resp, err := svc.PutObject(params)

		fmt.Printf("response %s", awsutil.StringValue(resp))

		imageName.NAME = uu

		imageNames = append(imageNames, imageName)
	}

	return imageNames
}

func PostImageToDBService(c *gin.Context, db *sql.DB) {
	dao.PostImageToDBDao(c, db)
	// var imageData util.ImageData
	// c.BindJSON(&imageData)

	// for _, imageName := range imageData.IMAGENAMES {
	// 	ins, err := db.Prepare("INSERT INTO images(article_uuid, image_name) VALUES(?,?)")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ins.Exec(imageData.ARTICLEUUID, imageName.NAME)
	// }
}
