# create user
GRANT ALL ON dbastore.* to 'dba'@'localhost' IDENTIFIED BY '123456';
GRANT ALL ON dbablog.* to 'dba'@'localhost' IDENTIFIED BY '123456';

-- Database dbastore
DROP DATABASE IF EXISTS `dbastore`;
CREATE DATABASE dbastore;
USE dbastore;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `user_id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(120) DEFAULT NULL,
  `last_name` varchar(120) DEFAULT NULL,
  `username` varchar(120) DEFAULT NULL,
  `email` varchar(80) DEFAULT NULL,
  `bio` varchar(2000) DEFAULT NULL,
  `age` int(2) DEFAULT NULL,
  `gender` varchar(8) DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT '0',
  `create_dttm` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_dttm` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`user_id`)
);

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `merchant` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `qty` int(9) unsigned DEFAULT NULL,
  `create_dttm` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_dttm` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`)
);

-- Database dbablog
DROP DATABASE IF EXISTS `dbablog`;
CREATE DATABASE dbablog;
USE dbablog;

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(8) DEFAULT NULL,
  `body` varchar(2000) DEFAULT NULL,
  `published` tinyint(1) DEFAULT '0',
  `create_dttm` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_dttm` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`)
);
