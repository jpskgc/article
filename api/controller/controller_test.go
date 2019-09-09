package controller

import (
	"article/api/service"
	"article/api/util"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockServiceInterface struct {
}

func (_m *MockServiceInterface) GetArticleService() []util.Article {
	return nil
}

func (_m *MockServiceInterface) GetSingleArticleService(c *gin.Context) util.Article {
	return util.Article{}
}

func (_m *MockServiceInterface) DeleteArticleService(c *gin.Context) {}

func (_m *MockServiceInterface) PostService(c *gin.Context) string {
	return ""
}

func (_m *MockServiceInterface) PostImageService(c *gin.Context) []util.ImageName {
	return nil
}

func (_m *MockServiceInterface) PostImageToDBService(c *gin.Context) {}

type ControllerSuite struct {
	suite.Suite
	controller *Controller
	service    service.ServiceInterface
}

func (s *ControllerSuite) SetupTest() {
	s.controller = NewController(s.service)
	s.controller.service = &MockServiceInterface{}
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) TestGetArticleController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.GetArticleController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}

func (s *ControllerSuite) TestGetSingleArticleController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.GetSingleArticleController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}

func (s *ControllerSuite) TestDeleteArticleController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.DeleteArticleController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}

func (s *ControllerSuite) TestPostController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.PostController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}

func (s *ControllerSuite) TestPostImageController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.PostImageController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}

func (s *ControllerSuite) TestPostImageToDBController() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	s.controller.PostImageToDBController(c)
	assert.Equal(s.T(), 200, c.Writer.Status())
}
