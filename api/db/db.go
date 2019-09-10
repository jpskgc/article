package db

import "database/sql"

type Database struct {
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	DATABASE       *sql.DB
}

func NewDatabase(user, password, host string) *Database {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/article")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS article;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use article;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `articles` (`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,uuid varchar(36), `title` VARCHAR(100) NOT NULL,`content` TEXT NOT NULL ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table IF NOT EXISTS images (id int AUTO_INCREMENT NOT NULL PRIMARY KEY, article_uuid varchar(36), image_name varchar(50));	")
	if err != nil {
		panic(err)
	}
	Database := new(Database)
	Database.DATABASE = db
	return Database
}
