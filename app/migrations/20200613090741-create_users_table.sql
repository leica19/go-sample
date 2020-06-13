
-- +migrate Up
CREATE TABLE `users` (
  `tenant_if` int(11) NOT NULL,
  `name` date NOT NULL,
  `age` varchar(14) NOT NULL,
  `sex` varchar(16) NOT NULL,
  `hire_date` date NOT NULL,
    PRIMARY KEY (emp_no)
)
-- +migrate Down
drop table `users`;