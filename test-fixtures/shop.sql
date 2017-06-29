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

-- INSERT - user
INSERT INTO `user` (`first_name`, `last_name`, `username`, `email`, `bio`, `age`, `gender`, `is_admin`)
VALUES ('James', 'Adu', 'j.adu', 'j.adu@google.com', 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim.', 23, 'male', 1);

INSERT INTO `user` (`first_name`, `last_name`, `username`, `email`, `bio`, `age`, `gender`, `is_admin`)
VALUES ('Florence', 'Park', 'fpark', 'fpark@yahoo.com', 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim.', 38, 'female', 1);

INSERT INTO `user` (`first_name`, `last_name`, `username`, `email`, `bio`, `age`, `gender`)
VALUES ('Tomi', 'Aya', 'ayatoms', 'aya.tomi@google.com', 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim.', 15, 'male');

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

-- INSERT - Product
INSERT INTO `product` (`name`, `brand`, `merchant`, `price`, `qty`)
VALUES ('Certified Refurbished All-New Echo Dot (2nd Generation) - Black', 'Amazon', 'Amazon', 44.99, 90000);

INSERT INTO `product` (`name`, `brand`, `merchant`, `price`, `qty`)
VALUES ('Lacoste Men Quartz Gold and Leather Automatic Watch, Color:Brown (Model: 2010871)', 'Lacoste', 'Amazon', 147.30, 45);

INSERT INTO `product` (`name`, `brand`, `merchant`, `price`, `qty`)
VALUES ('Dirt Devil Vacuum Cleaner Quick Lite Plus Bagless Corded Upright Vacuum UD20015', 'Dirt Devil', 'Amazon', 58.29, 300)
