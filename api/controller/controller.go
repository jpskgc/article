package controller

import (
	"database/sql"
	"net/http"

	"github.com/jpskgc/article/api/service"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
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

func GetArticleController(c *gin.Context, db *sql.DB) {
	articles := service.GetArticleService(c, db)
	// var articles []Article
	// results, err := db.Query("SELECT * FROM articles")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// article := Article{}
	// for results.Next() {
	// 	err = results.Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	} else {
	// 		articles = append(articles, article)
	// 	}
	// }
	c.JSON(http.StatusOK, articles)
}

func GetSingleArticleController(c *gin.Context, db *sql.DB) {
	article := service.GetSingleArticleService(c, db)

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
	// for rows.Next() {
	// 	imageName := ImageName{}
	// 	err := rows.Scan(&imageName.NAME)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	article.IMAGENAMES = append(article.IMAGENAMES, imageName)
	// }
	c.JSON(http.StatusOK, article)
}

func DeleteArticleController(c *gin.Context, db *sql.DB) {
	service.DeleteArticleService(c, db)
	// id := c.Params.ByName("id")
	// _, err := db.Exec("DELETE FROM articles WHERE id= ?", id)
	// if err != nil {
	// 	log.Fatalf("db.Exec(): %s\n", err)
	// }
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostController(c *gin.Context, db *sql.DB) {
	uu := service.PostService(c, db)
	// u, err := uuid.NewRandom()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// uu := u.String()
	// var article Article
	// c.BindJSON(&article)
	// ins, err := db.Prepare("INSERT INTO articles(uuid, title,content) VALUES(?,?,?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ins.Exec(uu, article.TITLE, article.CONTENT)

	c.JSON(http.StatusOK, gin.H{"uuid": uu})
}

func PostImageController(c *gin.Context) {

	imageNames := service.PostImageService(c)
	// creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	// cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
	// svc := s3.New(session.New(), cfg)

	// form, _ := c.MultipartForm()

	// files := form.File["images[]"]

	// var imageNames []ImageName
	// imageName := ImageName{}

	// for _, file := range files {

	// 	f, err := file.Open()

	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	defer f.Close()

	// 	size := file.Size
	// 	buffer := make([]byte, size)

	// 	u, err := uuid.NewRandom()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	uu := u.String()

	// 	f.Read(buffer)
	// 	fileBytes := bytes.NewReader(buffer)
	// 	fileType := http.DetectContentType(buffer)
	// 	path := "/media/" + uu
	// 	params := &s3.PutObjectInput{
	// 		Bucket:        aws.String("article-s3-jpskgc"),
	// 		Key:           aws.String(path),
	// 		Body:          fileBytes,
	// 		ContentLength: aws.Int64(size),
	// 		ContentType:   aws.String(fileType),
	// 	}
	// 	resp, err := svc.PutObject(params)

	// 	fmt.Printf("response %s", awsutil.StringValue(resp))

	// 	imageName.NAME = uu

	// 	imageNames = append(imageNames, imageName)
	// }

	c.JSON(http.StatusOK, imageNames)
}

func PostImageToDBController(c *gin.Context, db *sql.DB) {
	service.PostImageToDBService(c, db)
	// var imageData ImageData
	// c.BindJSON(&imageData)

	// for _, imageName := range imageData.IMAGENAMES {
	// 	ins, err := db.Prepare("INSERT INTO images(article_uuid, image_name) VALUES(?,?)")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ins.Exec(imageData.ARTICLEUUID, imageName.NAME)
	// }
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
