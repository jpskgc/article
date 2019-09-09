package service

import (
	"article/api/dao"
	"article/api/util"
	"bytes"
	"database/sql"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockDaoInterface struct {
}

func (_m *MockDaoInterface) GetArticleDao() *sql.Rows {
	db, mock, _ := sqlmock.New()
	mockRows := mock.NewRows([]string{"id", "uuid", "title", "content"}).
		AddRow(1, "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
		AddRow(2, "844bc620-7336-41a3-9cb4-552a0024ff1c", "test2", "test2")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")
	return rows
}

func (_m *MockDaoInterface) GetSingleArticleDao(id string) (util.Article, *sql.Rows) {

	article := util.Article{
		ID:      1,
		UUID:    "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c",
		TITLE:   "test",
		CONTENT: "test",
	}

	db, mock, _ := sqlmock.New()
	mockRows := mock.NewRows([]string{"image_name"}).
		AddRow("1a90696f-4fe7-48f5-81a5-ca72c129f4b0").
		AddRow("3d997272-468f-4b66-91db-00c39f0ef717")
	mock.ExpectQuery("select").WillReturnRows(mockRows)
	rows, _ := db.Query("select")

	return article, rows
}

func (_m *MockDaoInterface) DeleteArticleDao(id string) {
	return
}

func (_m *MockDaoInterface) PostDao(article util.Article, uu string) {
	return

}

func (_m *MockDaoInterface) PostImageToDBDao(imageData util.ImageData) {
	return

}

func (_m *MockDaoInterface) PostImageToS3Dao(file *multipart.FileHeader, imageName string) {
	return

}

type ServiceSuite struct {
	suite.Suite
	service *Service
	dao     dao.DaoInterface
}

func (s *ServiceSuite) SetupTest() {
	s.service = NewService(s.dao)
	s.service.dao = &MockDaoInterface{}
}

func (s *ServiceSuite) TestGetArticleService() {

	articles := s.service.GetArticleService()

	var expectedArticles []util.Article

	expectedArticle1 := util.Article{
		ID:      1,
		UUID:    "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c",
		TITLE:   "test",
		CONTENT: "test",
	}
	expectedArticles = append(expectedArticles, expectedArticle1)

	expectedArticle2 := util.Article{
		ID:      2,
		UUID:    "844bc620-7336-41a3-9cb4-552a0024ff1c",
		TITLE:   "test2",
		CONTENT: "test2",
	}
	expectedArticles = append(expectedArticles, expectedArticle2)

	assert.Equal(s.T(), expectedArticles, articles)
}

func (s *ServiceSuite) TestGetSingleArticleService() {

	param := gin.Param{"id", "1"}
	params := gin.Params{param}
	req, _ := http.NewRequest("GET", "/article/1", nil)
	var context *gin.Context
	context = &gin.Context{Request: req, Params: params}

	article := s.service.GetSingleArticleService(context)

	expectedArticle := util.Article{
		ID:      1,
		UUID:    "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c",
		TITLE:   "test",
		CONTENT: "test",
	}

	imageName1 := util.ImageName{
		NAME: "1a90696f-4fe7-48f5-81a5-ca72c129f4b0",
	}
	expectedArticle.IMAGENAMES = append(expectedArticle.IMAGENAMES, imageName1)

	imageName2 := util.ImageName{
		NAME: "3d997272-468f-4b66-91db-00c39f0ef717",
	}

	expectedArticle.IMAGENAMES = append(expectedArticle.IMAGENAMES, imageName2)

	assert.Equal(s.T(), expectedArticle, article)
}

func (s *ServiceSuite) TestPostImageService() {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	w, err := mw.CreateFormFile("images[]", "test")
	if assert.NoError(s.T(), err) {
		_, err = w.Write([]byte("images[]"))
		assert.NoError(s.T(), err)
	}
	mw.Close()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())

	imageNames := s.service.PostImageService(c)

	assert.NotEmpty(s.T(), imageNames)
}

func (s *ServiceSuite) TestPostImageToDBService() {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"articleUUID\":\"bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c\", \"imageNames\":[{\"name\": \"1925c2de071aff40eca2ac15524fe139-300x300.jpg\"}]}"))
	c.Request.Header.Add("Content-Type", "")

	var obj struct {
		ARTICLEUUID string           `json:"articleUUID"`
		IMAGENAMES  []util.ImageName `json:"imageNames"`
	}
	assert.NoError(s.T(), c.ShouldBindJSON(&obj))
	assert.Equal(s.T(), "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", obj.ARTICLEUUID)
	assert.Equal(s.T(), "1925c2de071aff40eca2ac15524fe139-300x300.jpg", obj.IMAGENAMES[0].NAME)
	assert.Equal(s.T(), 0, w.Body.Len())
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
