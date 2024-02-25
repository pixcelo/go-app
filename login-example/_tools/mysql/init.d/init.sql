CREATE TABLE `user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `password` VARCHAR(60) NOT NULL,
  `salt` VARCHAR(30) NOT NULL,
  `state` VARCHAR(8) NOT NULL,
  `activate_token` VARCHAR(8) NOT NULL,
  `updated_at` DATETIME(6) NOT NULL,
  `created_at` DATETIME(6) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX email_idx (email)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=100001;
