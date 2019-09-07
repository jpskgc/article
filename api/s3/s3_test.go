package s3

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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

	objectStore.PostImageToS3(c)

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

	imageNames := objectStore.PostImageToS3(c)

	objectStore.DeleteS3Image(imageNames)

}
