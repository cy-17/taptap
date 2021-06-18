/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : sanqi

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 19/06/2021 01:43:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS `game`;
CREATE TABLE `game` (
  `game_id` bigint NOT NULL AUTO_INCREMENT,
  `game_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `release_at` date NOT NULL,
  `author` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'unknown',
  `introduction` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `classification` int NOT NULL DEFAULT '0',
  `cover_image` text COLLATE utf8mb4_general_ci,
  `detail_images` text COLLATE utf8mb4_general_ci,
  `tags` text COLLATE utf8mb4_general_ci,
  `icon` text COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`game_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of game
-- ----------------------------
BEGIN;
INSERT INTO `game` VALUES (1, 'yuanshen', '2021-06-17', 'mihoyo', 'haowan', 1, 'http://guolin.tech/book.png', 'http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png', '休闲,运动,养成', 'http://guolin.tech/book.png');
INSERT INTO `game` VALUES (2, 'benghuai3', '2021-06-17', 'mihoyo', 'haowan1', 1, 'http://guolin.tech/book.png', 'http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png', '休闲,运动,养成', 'http://guolin.tech/book.png');
INSERT INTO `game` VALUES (3, 'benghuai2', '2021-06-17', 'mihoyo', 'haowan2', 1, 'http://guolin.tech/book.png', 'http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png,http://guolin.tech/book.png', '休闲,运动,养成', 'http://guolin.tech/book.png');
COMMIT;

-- ----------------------------
-- Table structure for game_comment
-- ----------------------------
DROP TABLE IF EXISTS `game_comment`;
CREATE TABLE `game_comment` (
  `comment_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `game_id` bigint NOT NULL,
  `replied_id` bigint DEFAULT NULL,
  `pid` bigint DEFAULT NULL,
  `content` text COLLATE utf8mb4_general_ci,
  `create_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `score` int DEFAULT NULL,
  PRIMARY KEY (`comment_id`),
  KEY `foreign_gcuid_userid` (`user_id`),
  KEY `foreign_gcgid_gameid` (`game_id`),
  CONSTRAINT `foreign_gcgid_gameid` FOREIGN KEY (`game_id`) REFERENCES `game` (`game_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `foreign_gcuid_userid` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of game_comment
-- ----------------------------
BEGIN;
INSERT INTO `game_comment` VALUES (3, 2, 1, 0, 0, '真好玩21', '2021-06-18 10:23:08', 1);
INSERT INTO `game_comment` VALUES (4, 2, 1, 0, 0, '真好玩231', '2021-06-18 10:23:09', 2);
INSERT INTO `game_comment` VALUES (5, 2, 1, 0, 0, '真好玩231', '2021-06-18 10:23:10', 3);
INSERT INTO `game_comment` VALUES (6, 2, 1, 0, 0, '真好玩', '2021-06-18 10:23:11', 4);
INSERT INTO `game_comment` VALUES (7, 2, 1, 2, 6, '真好玩', '2021-06-17 20:59:43', 5);
INSERT INTO `game_comment` VALUES (8, 2, 1, 2, 6, '真好玩', '2021-06-18 10:23:15', 4);
INSERT INTO `game_comment` VALUES (9, 2, 1, 2, 6, '真好玩', '2021-06-17 20:59:59', 5);
INSERT INTO `game_comment` VALUES (10, 2, 1, 2, 5, '真好玩', '2021-06-18 10:23:15', 3);
INSERT INTO `game_comment` VALUES (11, 2, 1, 2, 5, '真好玩', '2021-06-18 10:23:16', 2);
INSERT INTO `game_comment` VALUES (12, 2, 1, 2, 5, '真好玩', '2021-06-18 10:23:17', 1);
INSERT INTO `game_comment` VALUES (13, 2, 1, 2, 4, '真好玩', '2021-06-18 10:23:18', 2);
INSERT INTO `game_comment` VALUES (14, 2, 1, 2, 4, '真好玩', '2021-06-18 10:23:19', 3);
INSERT INTO `game_comment` VALUES (15, 2, 1, 2, 4, '真好玩', '2021-06-18 10:23:19', 4);
INSERT INTO `game_comment` VALUES (127, 2, 1, 0, 0, '真玩哈哈哈', '2021-06-18 14:29:02', 5);
INSERT INTO `game_comment` VALUES (128, 2, 1, 0, 0, '真玩哈哈哈1', '2021-06-18 14:29:21', 5);
INSERT INTO `game_comment` VALUES (129, 2, 1, 0, 0, '真玩哈哈哈12', '2021-06-18 14:29:22', 5);
INSERT INTO `game_comment` VALUES (130, 2, 1, 0, 0, '真玩哈哈哈123', '2021-06-18 14:29:23', 5);
INSERT INTO `game_comment` VALUES (131, 2, 1, 0, 0, '真玩哈哈哈1234', '2021-06-18 14:29:27', 5);
INSERT INTO `game_comment` VALUES (132, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:06', 5);
INSERT INTO `game_comment` VALUES (133, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:12', 5);
INSERT INTO `game_comment` VALUES (134, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:12', 5);
INSERT INTO `game_comment` VALUES (135, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:13', 5);
INSERT INTO `game_comment` VALUES (136, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:13', 5);
INSERT INTO `game_comment` VALUES (137, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:14', 5);
INSERT INTO `game_comment` VALUES (138, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:14', 5);
INSERT INTO `game_comment` VALUES (139, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:15', 5);
INSERT INTO `game_comment` VALUES (140, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:15', 5);
INSERT INTO `game_comment` VALUES (141, 2, 1, 2, 127, '真玩哈哈哈1234', '2021-06-18 14:37:16', 5);
INSERT INTO `game_comment` VALUES (142, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:25', 5);
INSERT INTO `game_comment` VALUES (143, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:25', 5);
INSERT INTO `game_comment` VALUES (144, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:26', 5);
INSERT INTO `game_comment` VALUES (145, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:26', 5);
INSERT INTO `game_comment` VALUES (146, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:27', 5);
INSERT INTO `game_comment` VALUES (147, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:27', 5);
INSERT INTO `game_comment` VALUES (148, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:28', 5);
INSERT INTO `game_comment` VALUES (149, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:29', 5);
INSERT INTO `game_comment` VALUES (150, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:29', 5);
INSERT INTO `game_comment` VALUES (151, 2, 1, 2, 128, '真玩哈哈哈1234', '2021-06-18 14:37:30', 5);
INSERT INTO `game_comment` VALUES (152, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:33', 5);
INSERT INTO `game_comment` VALUES (153, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:34', 5);
INSERT INTO `game_comment` VALUES (154, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:34', 5);
INSERT INTO `game_comment` VALUES (155, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:35', 5);
INSERT INTO `game_comment` VALUES (156, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:35', 5);
INSERT INTO `game_comment` VALUES (157, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:36', 5);
INSERT INTO `game_comment` VALUES (158, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:37', 5);
INSERT INTO `game_comment` VALUES (159, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:37', 5);
INSERT INTO `game_comment` VALUES (160, 2, 1, 2, 129, '真玩哈哈哈1234', '2021-06-18 14:37:38', 5);
INSERT INTO `game_comment` VALUES (162, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:42', 5);
INSERT INTO `game_comment` VALUES (163, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:42', 5);
INSERT INTO `game_comment` VALUES (164, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:43', 5);
INSERT INTO `game_comment` VALUES (165, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:43', 5);
INSERT INTO `game_comment` VALUES (166, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:44', 5);
INSERT INTO `game_comment` VALUES (167, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:44', 5);
INSERT INTO `game_comment` VALUES (168, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:45', 5);
INSERT INTO `game_comment` VALUES (169, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:45', 5);
INSERT INTO `game_comment` VALUES (170, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:46', 5);
INSERT INTO `game_comment` VALUES (171, 2, 1, 2, 130, '真玩哈哈哈1234', '2021-06-18 14:37:47', 5);
INSERT INTO `game_comment` VALUES (172, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:37:58', 5);
INSERT INTO `game_comment` VALUES (173, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:37:59', 5);
INSERT INTO `game_comment` VALUES (174, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:37:59', 5);
INSERT INTO `game_comment` VALUES (175, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:37:59', 5);
INSERT INTO `game_comment` VALUES (176, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:00', 5);
INSERT INTO `game_comment` VALUES (177, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:00', 5);
INSERT INTO `game_comment` VALUES (178, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:01', 5);
INSERT INTO `game_comment` VALUES (179, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:01', 5);
INSERT INTO `game_comment` VALUES (180, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:02', 5);
INSERT INTO `game_comment` VALUES (181, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 14:38:02', 5);
INSERT INTO `game_comment` VALUES (182, 2, 1, 2, 131, '真玩哈哈哈1234', '2021-06-18 15:01:53', 0);
INSERT INTO `game_comment` VALUES (184, 2, 1, 2, 4, '真好玩', '2021-06-19 01:28:04', 5);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `sex` int DEFAULT NULL,
  `nickname` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `avatar` text COLLATE utf8mb4_general_ci COMMENT '用户头像地址，考虑实现',
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `phone` text COLLATE utf8mb4_general_ci,
  `introduction` text COLLATE utf8mb4_general_ci,
  `location` text COLLATE utf8mb4_general_ci,
  `birthday` date DEFAULT NULL,
  `last_access_at` datetime NOT NULL,
  `create_at` datetime NOT NULL,
  `passport` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, '1', '1', '1', '1', '1', '1', '2021-06-02', '2021-06-16 16:09:45', '2021-06-19 16:09:49', '1');
INSERT INTO `user` VALUES (2, 2, 'mickeymouse', 'sdzxvgszf', '123456', '13682277102', 'asdasdgggggg', 'china', '2021-06-17', '2021-06-19 01:09:58', '2021-06-17 00:32:15', 'testsanqi123');
INSERT INTO `user` VALUES (3, NULL, 'memedada', NULL, '123456', NULL, NULL, NULL, NULL, '2021-06-17 00:32:46', '2021-06-17 00:32:46', 'testsanqi1234');
INSERT INTO `user` VALUES (4, NULL, 'medada', NULL, '123456', NULL, NULL, NULL, NULL, '2021-06-17 00:33:08', '2021-06-17 00:33:08', 'testsanqi4433');
INSERT INTO `user` VALUES (5, NULL, 'mededa', NULL, '123456', NULL, NULL, NULL, NULL, '2021-06-18 09:43:07', '2021-06-18 09:43:07', 'testsanqi9966');
INSERT INTO `user` VALUES (6, NULL, 'meddida', NULL, '$2a$10$44.tPQHoTnLPxLxcvfVIN.FASVOn6mpxfc.ekWPJQFZcbyPOZVMeO', NULL, NULL, NULL, NULL, '2021-06-18 09:46:34', '2021-06-18 09:46:34', 'testsanqi9996');
INSERT INTO `user` VALUES (7, NULL, 'meddidag', NULL, '$2a$10$Xc3AEHdhkVFN.OAHOvFC/.vrFpzEBmc5ddi2j0tDSH9P8le.ILLAG', NULL, NULL, NULL, NULL, '2021-06-18 09:57:43', '2021-06-18 09:57:43', 'testsanqi9999');
INSERT INTO `user` VALUES (8, NULL, 'chenyuewen', NULL, '$2a$10$1Wy4bypaDUTXET8LcKlzOOd9B0.HdGGSmhWFghSsZgaPyNG.Yn5a.', NULL, NULL, NULL, NULL, '2021-06-19 00:58:02', '2021-06-19 00:58:02', 'testzhuce');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
