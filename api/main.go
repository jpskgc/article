package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/joho/godotenv"
)

type Article struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	TITLE      string      `json:"title"`
	CONTENT    string      `json:"content"`
	IMAGENAMES []ImageName `json:"imageNames"`
}

type ImageName struct {
	NAME string `json:"name"`
}

type ImageData struct {
	ARTICLEUUID string      `json:"articleUUID"`
	IMAGENAMES  []ImageName `json:"imageNames"`
}

type Param struct {
	Bucket      string
	Key         string
	Expires     string
	ContentType string
}

func main() {
	//TODO production
	err := godotenv.Load()
	if err != nil {
		//TODO production
	}
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/article")
	//db, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/article")
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
			var articles []Article
			results, err := db.Query("SELECT * FROM articles")
			if err != nil {
				panic(err.Error())
			}
			article := Article{}
			for results.Next() {
				err = results.Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
				if err != nil {
					panic(err.Error())
				} else {
					articles = append(articles, article)
				}
			}
			c.JSON(http.StatusOK, articles)
		})
		api.GET("/article/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			article := Article{}
			errArticle := db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
			if errArticle != nil {
				panic(errArticle.Error())
			}
			//TODO nil
			rows, errImage := db.Query("SELECT image_name FROM images WHERE article_uuid  = ?", article.UUID)
			if errImage != nil {
				panic(errImage.Error())
			}
			for rows.Next() {
				imageName := ImageName{}
				err = rows.Scan(&imageName.NAME)
				article.IMAGENAMES = append(article.IMAGENAMES, imageName)
			}
			c.JSON(http.StatusOK, article)
		})
		api.POST("/post", func(c *gin.Context) {
			u, err := uuid.NewRandom()
			if err != nil {
				fmt.Println(err)
				return
			}
			uu := u.String()
			var article Article
			c.BindJSON(&article)
			ins, err := db.Prepare("INSERT INTO articles(uuid, title,content) VALUES(?,?,?)")
			if err != nil {
				log.Fatal(err)
			}
			ins.Exec(uu, article.TITLE, article.CONTENT)

			c.JSON(http.StatusOK, gin.H{"uuid": uu})

		})
		api.POST("/post/image", func(c *gin.Context) {
			// var creds *credentials.Credentials
			// var err error
			//creds := credentials.NewStaticCredentials(os.Getenv("awsAccessKeyID"), os.Getenv("awsSecretAccessKey"), "")
			creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

			// //local
			// creds = credentials.NewSharedCredentials("", "default")
			// //creds.Expire()
			// _, err = creds.Get()
			// //TODO production credentials
			// if err != nil {
			// 	creds = credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{})
			// 	_, err = creds.Get()
			// }

			cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
			svc := s3.New(session.New(), cfg)

			form, _ := c.MultipartForm()

			files := form.File["images[]"]

			var imageNames []ImageName
			imageName := ImageName{}

			for _, file := range files {

				f, err := file.Open()

				if err != nil {
					log.Println(err)
				}

				defer f.Close()

				size := file.Size
				buffer := make([]byte, size)

				f.Read(buffer)
				fileBytes := bytes.NewReader(buffer)
				fileType := http.DetectContentType(buffer)
				path := "/media/" + file.Filename
				params := &s3.PutObjectInput{
					Bucket:        aws.String("article-s3-jpskgc"),
					Key:           aws.String(path),
					Body:          fileBytes,
					ContentLength: aws.Int64(size),
					ContentType:   aws.String(fileType),
				}
				resp, err := svc.PutObject(params)

				fmt.Printf("response %s", awsutil.StringValue(resp))

				imageName.NAME = file.Filename

				imageNames = append(imageNames, imageName)
			}

			c.JSON(http.StatusOK, imageNames)
		})
		api.POST("/post/image/db", func(c *gin.Context) {
			var imageData ImageData
			c.BindJSON(&imageData)

			for _, imageName := range imageData.IMAGENAMES {
				ins, err := db.Prepare("INSERT INTO images(article_uuid, image_name) VALUES(?,?)")
				if err != nil {
					log.Fatal(err)
				}
				ins.Exec(imageData.ARTICLEUUID, imageName.NAME)
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	router.Run(":2345")
}
