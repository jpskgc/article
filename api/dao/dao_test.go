package dao

import (
	"article/api/s3"
	"article/api/util"
	"database/sql"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/gin-gonic/gin"
)

type DaoSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	dao  *Dao
	s3   *s3.S3
}

func (s *DaoSuite) SetupTest() {

	var err error
	s.db, s.mock, err = sqlmock.New()
	s.Require().NoError(err)
	s.dao = NewDao(s.db, s.s3)
}

func (s *DaoSuite) TestGetArticleDao() {

	rows := s.mock.NewRows([]string{"id", "uuid", "title", "content"}).
		AddRow(1, "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
		AddRow(2, "844bc620-7336-41a3-9cb4-552a0024ff1c", "test2", "test2")

	s.mock.ExpectQuery("^SELECT (.+) FROM articles*").
		WillReturnRows(rows)

	results := s.dao.GetArticleDao()
	var articles []util.Article
	article := util.Article{}
	for results.Next() {
		err := results.Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
		if err != nil {
			panic(err.Error())
		} else {
			articles = append(articles, article)
		}
	}

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

func (s *DaoSuite) TestGetSingleArticleDao() {

	articleMockRows := s.mock.NewRows([]string{"id", "uuid", "title", "content"}).
		AddRow(1, "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test")

	s.mock.ExpectQuery("^SELECT (.+) FROM articles*").
		WithArgs("1").
		WillReturnRows(articleMockRows)

	imageMockRows := s.mock.NewRows([]string{"image_name"}).
		AddRow("1a90696f-4fe7-48f5-81a5-ca72c129f4b0").
		AddRow("3d997272-468f-4b66-91db-00c39f0ef717")

	s.mock.ExpectQuery("^SELECT (.+) FROM images*").
		WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c").
		WillReturnRows(imageMockRows)

	param := gin.Param{"id", "1"}
	params := gin.Params{param}
	req, _ := http.NewRequest("GET", "/article/1", nil)
	var context *gin.Context
	context = &gin.Context{Request: req, Params: params}

	article, imageRows := s.dao.GetSingleArticleDao(context)

	for imageRows.Next() {
		imageName := util.ImageName{}
		err := imageRows.Scan(&imageName.NAME)
		if err != nil {
			panic(err.Error())
		}
		article.IMAGENAMES = append(article.IMAGENAMES, imageName)
	}

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

func (s *DaoSuite) TestDeleteArticleDao() {

	articleMockRows := s.mock.NewRows([]string{"id", "uuid", "title", "content"}).
		AddRow(1, "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test")

	s.mock.ExpectQuery("^SELECT (.+) FROM articles*").
		WithArgs("1").
		WillReturnRows(articleMockRows)

	imageMockRows := s.mock.NewRows([]string{"image_name"}).
		AddRow("1a90696f-4fe7-48f5-81a5-ca72c129f4b0")

	s.mock.ExpectQuery("^SELECT (.+) FROM images*").
		WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c").
		WillReturnRows(imageMockRows)

	s.mock.ExpectBegin()

	s.mock.ExpectExec("^DELETE FROM articles*").WithArgs("1").
		WillReturnResult(sqlmock.NewErrorResult(nil))

	s.mock.ExpectExec("^DELETE FROM images*").WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c").
		WillReturnResult(sqlmock.NewErrorResult(nil))

	s.mock.ExpectCommit()

	s.dao.DeleteArticleDao("1")

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expections: %s", err)
	}
}

func (s *DaoSuite) TestPostDao() {

	prep := s.mock.ExpectPrepare("^INSERT INTO articles*")

	prep.ExpectExec().
		WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
		WillReturnResult(sqlmock.NewResult(1, 1))

	article := util.Article{
		ID:      1,
		TITLE:   "test",
		CONTENT: "test",
	}

	s.dao.PostDao(article, "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c")

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expections: %s", err)
	}
}

func (s *DaoSuite) TestPostImageToDBDao() {

	prep := s.mock.ExpectPrepare("^INSERT INTO images*")

	prep.ExpectExec().
		WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "b8119536-fad5-4ffa-ab71-2f96cca19697").
		WithArgs("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "118f4bd4-a477-4ea1-a90e-2257b69a6989").
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedImageData := util.ImageData{
		ARTICLEUUID: "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c",
	}

	var expectedImageName []util.ImageName

	expectedImageName1 := util.ImageName{
		NAME: "b8119536-fad5-4ffa-ab71-2f96cca19697",
	}

	expectedImageName = append(expectedImageName, expectedImageName1)

	expectedImageName2 := util.ImageName{
		NAME: "118f4bd4-a477-4ea1-a90e-2257b69a6989",
	}
	expectedImageName = append(expectedImageName, expectedImageName2)

	expectedImageData.IMAGENAMES = expectedImageName

	s.dao.PostImageToDBDao(expectedImageData)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expections: %s", err)
	}
}

func (s *DaoSuite) TearDownTest() {
	s.db.Close()
	s.Assert().NoError(s.mock.ExpectationsWereMet())
}

func TestDaoSuite(t *testing.T) {
	suite.Run(t, new(DaoSuite))
}
