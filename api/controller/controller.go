package controller

import (
	"article/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{service: service}
}

func (controller Controller) GetArticleController(c *gin.Context) {
	articles := controller.service.GetArticleService()
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

func (controller Controller) PostImageController(c *gin.Context) {
	imageNames := controller.service.PostImageService(c)
	c.JSON(http.StatusOK, imageNames)
}

func (controller Controller) PostImageToDBController(c *gin.Context) {
	controller.service.PostImageToDBService(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
