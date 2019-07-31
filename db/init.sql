CREATE DATABASE IF NOT EXISTS article;
use article;
CREATE TABLE IF NOT EXISTS `articles` (`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",uuid varchar(36), `title` VARCHAR(100) NOT NULL COMMENT "タイトル",`content` TEXT NOT NULL COMMENT "本文" ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
create table IF NOT EXISTS images (id int AUTO_INCREMENT NOT NULL PRIMARY KEY, article_uuid varchar(36), image_name varchar(50));
