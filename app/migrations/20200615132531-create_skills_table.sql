
-- +migrate Up
CREATE TABLE `skills` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `golang` tinyint(1) DEFAULT NULL,
  `docker` tinyint(1) DEFAULT NULL,
  `rust` tinyint(1) DEFAULT NULL,
  `php` tinyint(1) DEFAULT NULL,
  `javascript` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_skills_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
-- +migrate Down
drop table `skills`;