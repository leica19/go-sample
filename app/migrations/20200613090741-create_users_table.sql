
-- +migrate Up
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `tenant_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `age` int(10) unsigned DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `skill_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
-- +migrate Down
drop table `users`;