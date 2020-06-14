
-- +migrate Up
CREATE TABLE `users` (
  `tenant_id` int(11) NOT NULL,
  `name` VARCHAR(30) NOT NULL,
  `age` int(3) NOT NULL,
  `sex` varchar(5) NOT NULL
);
-- +migrate Down
drop table `users`;