-- mysql> show columns from account;
-- +---------------+--------------+------+-----+-------------------+----------------+
-- | Field         | Type         | Null | Key | Default           | Extra          |
-- +---------------+--------------+------+-----+-------------------+----------------+
-- | id            | bigint(20)   | NO   | PRI | NULL              | auto_increment |
-- | username      | varchar(255) | NO   | UNI | NULL              |                |
-- | password_hash | varchar(255) | NO   |     | NULL              |                |
-- | display_name  | varchar(255) | YES  |     | NULL              |                |
-- | avatar        | text         | YES  |     | NULL              |                |
-- | header        | text         | YES  |     | NULL              |                |
-- | note          | text         | YES  |     | NULL              |                |
-- | create_at     | datetime     | NO   |     | CURRENT_TIMESTAMP |                |
-- +---------------+--------------+------+-----+-------------------+----------------+
CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- mysql> show columns from status;
-- +------------+------------+------+-----+-------------------+----------------+
-- | Field      | Type       | Null | Key | Default           | Extra          |
-- +------------+------------+------+-----+-------------------+----------------+
-- | id         | bigint(20) | NO   | PRI | NULL              | auto_increment |
-- | account_id | bigint(20) | NO   | MUL | NULL              |                |
-- | content    | text       | NO   |     | NULL              |                |
-- | create_at  | datetime   | NO   |     | CURRENT_TIMESTAMP |                |
-- +------------+------------+------+-----+-------------------+----------------+
CREATE TABLE `status` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) NOT NULL,
  `content` text NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_account_id` (`account_id`),
  CONSTRAINT `fk_status_account_id` FOREIGN KEY (`account_id`) REFERENCES  `account` (`id`)
);
