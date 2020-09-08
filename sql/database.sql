/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 127.0.0.1:3306
 Source Schema         : ppf

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 08/09/2020 16:14:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat_user
-- ----------------------------
DROP TABLE IF EXISTS `chat_user`;
CREATE TABLE `chat_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `send_uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送用户id',
  `recv_uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '接收用户id',
  `message_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '消息类型，1文字',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `created_at` int(10) NULL DEFAULT NULL COMMENT '发送时间',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态，0未读，1已读，-1撤回',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '私聊聊天记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for friend_add
-- ----------------------------
DROP TABLE IF EXISTS `friend_add`;
CREATE TABLE `friend_add`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `f_uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友uid',
  `channel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '加好友途径，搜索、附近的人',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '添加原因',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态，1同意，0等待，-1拒绝',
  `request_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加好友请求时间戳',
  `pass_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '通过好友时间戳',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '添加好友表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for friend_list
-- ----------------------------
DROP TABLE IF EXISTS `friend_list`;
CREATE TABLE `friend_list`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `f_uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '好友uid',
  `channel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '加好友途径，搜索、附近的人',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '添加原因',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT '1加的那一方，2被加的那一方',
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique`(`uid`, `f_uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `password_salt` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码盐值',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `country` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '国家',
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '省份',
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `sex` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别，0未知，1男，2女',
  `real_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `id_card` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `longitude` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '经度',
  `latitude` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '纬度',
  `last_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录/注册ip',
  `register_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间戳',
  `login_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
