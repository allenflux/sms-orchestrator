/*
 Navicat Premium Dump SQL

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80040 (8.0.40)
 Source Host           : 127.0.0.1:33062
 Source Schema         : sms

 Target Server Type    : MySQL
 Target Server Version : 80040 (8.0.40)
 File Encoding         : 65001

 Date: 28/12/2024 15:55:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for device_list
-- ----------------------------
DROP TABLE IF EXISTS `device_list`;
CREATE TABLE `device_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_id` varchar(255) NOT NULL COMMENT '设备id',
  `number` varchar(255) DEFAULT NULL COMMENT '设备号码',
  `active_time` timestamp NULL DEFAULT NULL COMMENT '激活时间',
  `owner_account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '子账号',
  `assigned_items` varchar(255) NOT NULL COMMENT '所属项目',
  `sent_status` tinyint NOT NULL COMMENT '发送状态，1-异常 2-占用 3-空闲',
  `quantity_sent` int NOT NULL COMMENT '发送数量',
  `quantity_all` int NOT NULL COMMENT '全部数量',
  `device_status` tinyint NOT NULL COMMENT '设备状态，1-异常 2-正常',
  `group_name` varchar(255) DEFAULT NULL COMMENT '分组名称',
  `group_id` int DEFAULT NULL COMMENT '分组id',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `device_number` varchar(255) DEFAULT NULL COMMENT '设备序列号',
  `start_at` timestamp NULL DEFAULT NULL COMMENT '发送时间',
  `assigned_items_id` int DEFAULT NULL COMMENT '项目id',
  `owner_account_id` int NOT NULL COMMENT '子账号id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '操作者的用户id',
  `user_name` varchar(255) NOT NULL COMMENT '操作者的用户名',
  `client_ip` varchar(255) NOT NULL COMMENT '操作ip',
  `function` varchar(255) NOT NULL COMMENT '操作功能',
  `note` varchar(255) NOT NULL COMMENT '操作内容',
  `system_id` int NOT NULL COMMENT '后台id',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`) USING BTREE COMMENT '加速根据时间筛选',
  KEY `idx_user_id` (`user_id`) COMMENT '加速根据用户筛选'
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL COMMENT '功能父id',
  `name` varchar(255) DEFAULT NULL COMMENT '规则名称',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '功能名称',
  `icon` varchar(255) DEFAULT NULL COMMENT 'icon图标',
  `menu_type` tinyint DEFAULT NULL COMMENT '菜单类型，1-目录 2-菜单 3-按钮',
  `is_hide` tinyint DEFAULT NULL COMMENT '显示状态，1-隐藏 2-显示',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '路由地址',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `is_cached` tinyint DEFAULT NULL COMMENT '是否缓存，1-不缓存，2-缓存',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for project_list
-- ----------------------------
DROP TABLE IF EXISTS `project_list`;
CREATE TABLE `project_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',
  `quantity_device` int NOT NULL COMMENT '设备数量',
  `associated_account` varchar(255) DEFAULT NULL COMMENT '关联账号',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `associated_account_id` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL COMMENT '角色id',
  `permission_id` int NOT NULL COMMENT '权限id',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sms_chart_log
-- ----------------------------
DROP TABLE IF EXISTS `sms_chart_log`;
CREATE TABLE `sms_chart_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_name` varchar(255) NOT NULL COMMENT '项目名称',
  `project_id` int NOT NULL COMMENT '项目id',
  `target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',
  `device_number` varchar(255) NOT NULL COMMENT '执行设备号',
  `sms_topic` varchar(255) NOT NULL COMMENT '短信主题',
  `sms_content` varchar(255) NOT NULL COMMENT '短信内容',
  `sms_status` tinyint NOT NULL COMMENT '短信发送状态，1-失败 2-成',
  `associated_account` varchar(255) NOT NULL COMMENT '所属子账号',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `sent_or_receive` int NOT NULL COMMENT '1表示此条短信是发送 2表示此条短信是接收',
  `associated_account_id` int NOT NULL COMMENT '子账号id',
  `row_hash` varchar(255) NOT NULL COMMENT 'log hash 防止重复上报',
  `task_id` int NOT NULL COMMENT 'Mission Task的id 任务 id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `rowhash` (`row_hash`),
  KEY `device` (`device_number`) USING BTREE,
  KEY `sub_user_id` (`associated_account_id`) USING BTREE,
  KEY `target_phone_number` (`target_phone_number`) USING BTREE,
  KEY `project_id` (`project_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1081310 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sms_mission_record
-- ----------------------------
DROP TABLE IF EXISTS `sms_mission_record`;
CREATE TABLE `sms_mission_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `task_name` varchar(255) NOT NULL COMMENT '任务名称',
  `sub_task_id` int NOT NULL COMMENT '子任务id',
  `target_phone_number` varchar(255) NOT NULL COMMENT '目标手机号',
  `device_number` varchar(255) NOT NULL COMMENT '执行设备号',
  `sms_topic` varchar(255) NOT NULL COMMENT '短信主题',
  `sms_content` varchar(255) NOT NULL COMMENT '短信内容',
  `sms_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '短信发送状态，2-失败 1-成功',
  `associated_account` varchar(255) NOT NULL COMMENT '所属子账号',
  `project_name` varchar(255) NOT NULL COMMENT '所属项目名称',
  `project_id` int NOT NULL COMMENT '所属项目id',
  `start_time` timestamp NOT NULL COMMENT '开始时间',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `associated_account_id` int NOT NULL COMMENT '子账户id',
  `row_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '每行内容的hash串 为了防止同样的记录重复提交',
  `reason` varchar(255) DEFAULT NULL COMMENT '失败原因',
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash` (`row_hash`),
  KEY `target_phone` (`target_phone_number`),
  KEY `device_number` (`device_number`)
) ENGINE=InnoDB AUTO_INCREMENT=63380 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sms_mission_report
-- ----------------------------
DROP TABLE IF EXISTS `sms_mission_report`;
CREATE TABLE `sms_mission_report` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_id` int NOT NULL COMMENT '项目id',
  `task_name` varchar(255) NOT NULL COMMENT '任务名称',
  `file_name` varchar(255) NOT NULL COMMENT '文件名',
  `device_quota` varchar(255) NOT NULL COMMENT '执行设备',
  `task_status` tinyint NOT NULL DEFAULT '1' COMMENT '任务状态，1-待发送 2-发送中 3-已发送 4-已撤销',
  `sms_quantity` int NOT NULL COMMENT '短信总条数',
  `surplus_quantity` int NOT NULL COMMENT '剩余数量',
  `quantity_sent` int NOT NULL DEFAULT '0' COMMENT '已发送数量',
  `associated_account` varchar(255) NOT NULL COMMENT '所属子账号',
  `interval_time` varchar(255) NOT NULL COMMENT '间隔时间(秒)',
  `start_time` timestamp NOT NULL COMMENT '任务开始时间',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',
  `associated_account_id` int DEFAULT NULL COMMENT '所属子账号id',
  `group_id` int DEFAULT NULL COMMENT '分组id',
  `sent_success_quantity` int NOT NULL DEFAULT '0' COMMENT '发送成功数量',
  `sent_fail_quantity` int NOT NULL COMMENT '发送失败数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sms_trace_task
-- ----------------------------
DROP TABLE IF EXISTS `sms_trace_task`;
CREATE TABLE `sms_trace_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `target_number` varchar(255) DEFAULT NULL,
  `device_number` varchar(255) DEFAULT NULL,
  `task_id` int NOT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  `delete_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40166 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sub_group
-- ----------------------------
DROP TABLE IF EXISTS `sub_group`;
CREATE TABLE `sub_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sub_user_id` int NOT NULL COMMENT '子账号id',
  `project_id` int NOT NULL COMMENT '项目id',
  `sub_group_name` varchar(255) NOT NULL COMMENT '分组名称',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '修改时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `salt` varchar(255) NOT NULL COMMENT '密码盐',
  `status` tinyint NOT NULL COMMENT '账号状态，1-停用 2-启用',
  `role_id` int NOT NULL COMMENT '角色权限id',
  `system_id` int NOT NULL COMMENT '所属后台id',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `update_at` timestamp NOT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=890 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
