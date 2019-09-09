package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"article/api/dao"

	"article/api/util"
)

type Service struct {
	dao dao.DaoInterface
}

func NewService(dao dao.DaoInterface) *Service {
	return &Service{dao: dao}
}

func (s Service) GetArticleService() []util.Article {
	var articles []util.Article

	results := s.dao.GetArticleDao()

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

func (s Service) GetSingleArticleService(c *gin.Context) util.Article {

	article, rows := s.dao.GetSingleArticleDao(c)

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

func (s Service) DeleteArticleService(c *gin.Context) {
	id := c.Params.ByName("id")
	s.dao.DeleteArticleDao(id)
}

func (s Service) PostService(c *gin.Context) string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	uu := u.String()
	var article util.Article
	c.BindJSON(&article)
	s.dao.PostDao(article, uu)

	return uu
}

func (s Service) PostImageService(c *gin.Context) []util.ImageName {

	form, _ := c.MultipartForm()
	files := form.File["images[]"]

	var imageNames []util.ImageName
	imageName := util.ImageName{}

	for _, file := range files {

		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()

		s.dao.PostImageToS3Dao(file, uu)

		imageName.NAME = uu

		imageNames = append(imageNames, imageName)
	}

	return imageNames
}

func (s Service) PostImageToDBService(c *gin.Context) {
	var imageData util.ImageData
	c.BindJSON(&imageData)
	s.dao.PostImageToDBDao(imageData)
}
