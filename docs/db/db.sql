
CREATE DATABASE IF NOT EXISTS some;
USE some;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '昵称',
  `phone` bigint(20) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wealth_cost_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '1',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parrent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父分类ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='支出分类';


CREATE TABLE `wealth_cost_details` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '1',
  `category_id` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '分类',
  `occur_date` date NOT NULL DEFAULT '1900-01-01',
  `content` varchar(45) NOT NULL DEFAULT '' COMMENT '摘要',
  `amount` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '金额',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '说明',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_occur_date` (`occur_date`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='支出明细';

CREATE TABLE `wealth_cost_statistics_months` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `year` int(10) unsigned NOT NULL COMMENT '年份',
  `year_month` int(10) unsigned NOT NULL COMMENT '年月',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别',
  `total` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '合计',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `wealth_cost_statistics_weeks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `year` int(10) unsigned NOT NULL,
  `start_date` date NOT NULL DEFAULT '2000-01-01' COMMENT '周数',
  `end_date` date NOT NULL DEFAULT '2000-01-01' COMMENT '结束日期',
  `category_id` int(10) DEFAULT '0' COMMENT '类别',
  `total` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '合计',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
