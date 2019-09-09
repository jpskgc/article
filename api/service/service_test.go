package service

import (
	"article/api/dao"
	"article/api/util"
	"database/sql"
	"mime/multipart"
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

func (_m *MockDaoInterface) GetSingleArticleDao(c *gin.Context) (util.Article, *sql.Rows) {
	return util.Article{}, nil
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

func (s *ServiceSuite) TestGetArticleService(t *testing.T) {

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

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
