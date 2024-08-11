/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : 127.0.0.1:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 06/08/2024 18:33:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `userid` smallint NULL DEFAULT NULL COMMENT '用户ID',
  `goodsid` smallint NULL DEFAULT NULL COMMENT '商品ID',
  `goods_count` smallint NULL DEFAULT NULL COMMENT '下单商品数量',
  `goods_price` float NULL DEFAULT NULL COMMENT '总价格',
  `pay_type` tinyint NULL DEFAULT NULL COMMENT '支付方式',
  `state` smallint NULL DEFAULT NULL COMMENT '支付状态',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `order_numbering` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
