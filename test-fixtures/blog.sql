-- Database dbablog
DROP DATABASE IF EXISTS `dbablog`;
CREATE DATABASE dbablog;
USE dbablog;

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `body` text DEFAULT NULL,
  `published` tinyint(1) DEFAULT '0',
  `create_dttm` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_dttm` timestamp ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`)
);

-- INSERT - Post

INSERT INTO `post` (title, body, published)
VALUES ('March ASG meeting – Saturday 18th March', 'Dr Bernard Asabere, the line manager of the Radio Astronomy Observatory will be speaking on “Radio Astronomy in Ghana” for the March meeting of the ASG (Astronomical Society of Ghana) at the Planetarium. All are welcome, so join us to find out about Ghana’s ground-breaking astronomy project.  Full details to follow soon!', 1);

INSERT INTO `post` (title, body, published)
VALUES ('Product updates | March 8, 2017', 'The database service that Google uses for its own mission-critical applications is now publicly available. This fully managed DBaaS offers petabyte-scale distributed transactions, high availability, and fast, global access to dataThe database service that Google uses for its own mission-critical applications is now publicly available. This fully managed DBaaS offers petabyte-scale distributed transactions, high availability, and fast, global access to data', 1);

INSERT INTO `post` (title, body, published)
VALUES ('Exciting Announcements From Intersect 2017', 'There were some seriously exciting announcements at Intersect 2017 tod‍ay! New Nanodegree programs, new hiring partners, big prize money challenges, and more. Read on to learn everything you need to know about what was announced, and all the great new ways you have to keep learning for the jobs of to‍day, tom‍orrow, and beyond!', 0)
