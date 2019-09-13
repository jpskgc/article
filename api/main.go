package main

import (
	"article/api/dao"
	"article/api/db"
	"article/api/handler"
	"article/api/s3"
	"article/api/service"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"article/api/controller"
)

func main() {

	log.Println("MYSQL_USER: " + os.Getenv("MYSQL_USER"))

	db := db.NewDatabase(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	s3 := s3.NewS3(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	dao := dao.NewDao(db.DATABASE, s3)
	service := service.NewService(dao)
	cntlr := controller.NewController(service)
	router := handler.NewHandler(cntlr)

	router.Router.Run(":2345")
}
