package controller

import (
	"net/http"

	"article/api/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func (controller Controller) GetArticleController(c *gin.Context) {
	articles := controller.service.GetArticleService(c)
	c.JSON(http.StatusOK, articles)
}

func (controller Controller) GetSingleArticleController(c *gin.Context) {
	article := controller.service.GetSingleArticleService(c)

	c.JSON(http.StatusOK, article)
}

func (controller Controller) DeleteArticleController(c *gin.Context) {
	controller.service.DeleteArticleService(c)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (controller Controller) PostController(c *gin.Context) {
	uu := controller.service.PostService(c)

	c.JSON(http.StatusOK, gin.H{"uuid": uu})
}

func PostImageController(c *gin.Context) {

	imageNames := service.PostImageService(c)

	c.JSON(http.StatusOK, imageNames)
}

func (controller Controller) PostImageToDBController(c *gin.Context) {
	controller.service.PostImageToDBService(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}
