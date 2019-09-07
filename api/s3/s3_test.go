package s3

import (
	"article/api/util"
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var objectStore *S3

func prepareObjectStore() {

	var (
		appid  = os.Getenv("AWS_ACCESS_KEY_ID")
		secret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	)

	objectStore = NewS3(appid,
		secret,
	)
}

func TestPostImageToS3(t *testing.T) {
	prepareObjectStore()

	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	w, err := mw.CreateFormFile("file", "test")
	if assert.NoError(t, err) {
		_, err = w.Write([]byte("test"))
		assert.NoError(t, err)
	}
	mw.Close()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())

	form, _ := c.MultipartForm()
	files := form.File["test"]

	for _, file := range files {
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()
		err = objectStore.PostImageToS3(file, uu)
		if err != nil {
			t.Fatalf("PostImageToS3 error %s", err)
		}
	}
}

func TestDeleteS3Image(t *testing.T) {
	prepareObjectStore()

	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	w, err := mw.CreateFormFile("file", "test")
	if assert.NoError(t, err) {
		_, err = w.Write([]byte("test"))
		assert.NoError(t, err)
	}
	mw.Close()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", buf)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())

	form, _ := c.MultipartForm()
	files := form.File["test"]

	for _, file := range files {
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()
		err = objectStore.PostImageToS3(file, uu)
		if err != nil {
			t.Fatalf("PostImageToS3 error %s", err)
		}

		expectedImageName := util.ImageName{
			NAME: uu,
		}
		err = objectStore.DeleteS3Image(expectedImageName)
		if err != nil {
			t.Fatalf("DeleteS3Image error %s", err)
		}
	}

}
