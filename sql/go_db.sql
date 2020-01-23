/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50527
Source Host           : localhost:3306
Source Database       : go_db

Target Server Type    : MYSQL
Target Server Version : 50527
File Encoding         : 65001

Date: 2020-01-23 17:01:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for artikel
-- ----------------------------
DROP TABLE IF EXISTS `artikel`;
CREATE TABLE `artikel` (
  `id_artikel` int(11) NOT NULL AUTO_INCREMENT,
  `nama_artikel` varchar(100) DEFAULT NULL,
  `keterangan` text,
  `status` int(11) DEFAULT '0',
  PRIMARY KEY (`id_artikel`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of artikel
-- ----------------------------
INSERT INTO `artikel` VALUES ('6', 'ELECTRIC POWER GENERATION', 'To meet the rapidly increasing need for electrical power, Trakindo provides a comprehensive range of solutions to customer needs. We are committed to provide solution for all of your electrical and power needs.  Sustained rapid growth in the manufacturing, telecommunications and other industrial sectors in Indonesia in recent years has led to an increasing demand for reliable power supplies. In many areas around the country, the existing electrical supply and distribution networks are unable to cope with the increasing demand.  In order to ensure that electric power availability and productivity satisfy the benchmarks set by our customers, top-class maintenance management is critical. That’s why we forge close working relationships and partnerships with our customers to ensure that their resources and skills are fully harnessed so as to maximize machinery availability and keep costs down. Our maintenance and repair programs are tailored to suit the requirements of specific projects and range from basic parts supply to fixed cost repair and maintenance contracts. Trakindo is your proactive and consultative partner to advance your power solution.', null);

-- ----------------------------
-- Table structure for kontak
-- ----------------------------
DROP TABLE IF EXISTS `kontak`;
CREATE TABLE `kontak` (
  `id_kontak` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `keterangan` text,
  PRIMARY KEY (`id_kontak`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of kontak
-- ----------------------------

-- ----------------------------
-- Table structure for names
-- ----------------------------
DROP TABLE IF EXISTS `names`;
CREATE TABLE `names` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `email` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of names
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'fahri_nugraha76@yahoo.co.id', 'fachri', 'test', '$2a$10$9X2NVLXanv1sxYUQ1DVGeOx2O.1cCBYnwY96hFpZtSJNjrRll7ppC');
INSERT INTO `users` VALUES ('2', 'fachri.nugraha@trakindo.co.id', 'test', 'test', '$2a$10$Fy6mUyP2wNP1gGgs0myE/e0ZzdwZjGk3aPi3U9RIp7QTserRqJ9Da');
