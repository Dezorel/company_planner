CREATE DATABASE IF NOT EXISTS Company_planner;

    -- ----------------------------
-- Table structure for cabinet_size
-- ----------------------------
DROP TABLE IF EXISTS `Cabinet_size`;
CREATE TABLE `Cabinet_size`  (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `cabinet_size` int(11) NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 INDEX `size`(`cabinet_size`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `Companies`;
CREATE TABLE `Companies`  (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `company_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                              `date_time_created` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                              PRIMARY KEY (`id`) USING BTREE,
                              INDEX `company_name`(`company_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for cabinets
-- ---------------------------
DROP TABLE IF EXISTS `Cabinets`;
CREATE TABLE `Cabinets` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `number` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                            `company_id` int(11) NOT NULL,
                            `specifications` text NULL DEFAULT NULL,
                            `size_id` int(11) NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `size_id`(`size_id`) USING BTREE,
                            INDEX `company_id`(`company_id`) USING BTREE,
                            INDEX `number`(`number`) USING BTREE,
                            CONSTRAINT `size_id` FOREIGN KEY (`size_id`) REFERENCES `Cabinet_size` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                            CONSTRAINT `company_id` FOREIGN KEY (`company_id`) REFERENCES `Companies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cabinets schedule
-- ----------------------------

DROP TABLE IF EXISTS `Cabinets_schedule`;
CREATE TABLE `Cabinets_schedule`  (
                              `cabinet_id` int(11) NOT NULL,
                              `date_time_start` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
                              `date_time_end` timestamp(0) NULL DEFAULT NULL,
                              INDEX `cabinet_id`(`cabinet_id`) USING BTREE,
                              CONSTRAINT `cabinet_id` FOREIGN KEY (`cabinet_id`) REFERENCES `Cabinets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;