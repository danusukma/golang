/*
SQLyog Community v12.2.2 (64 bit)
MySQL - 5.7.31 : Database - sensors_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`sensors_db` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `sensors_db`;

/*Table structure for table `sensors` */

DROP TABLE IF EXISTS `sensors`;

CREATE TABLE `sensors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensorvalue` int(11) DEFAULT NULL,
  `id1` int(11) DEFAULT NULL,
  `id2` varchar(10) DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;

/*Data for the table `sensors` */

insert  into `sensors`(`id`,`sensorvalue`,`id1`,`id2`,`timestamp`) values 
(6,64,1,'A','2022-06-04 13:21:16'),
(7,49,2,'B','2022-06-04 13:22:08'),
(8,33,3,'A','2022-06-04 14:07:37'),
(11,44,3,'A','2022-06-04 14:37:19'),
(10,55,1,'C','2022-06-04 14:10:41'),
(12,46,3,'A','2022-06-04 15:43:47');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
