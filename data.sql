CREATE TABLE users (
  `id`         CHAR(36) NOT NULL,
  `email`      VARCHAR(50) NOT NULL,
  `password`   VARCHAR(50) NOT NULL,
  `salt`       VARCHAR(50) NULL,
  `username`  VARCHAR(50) NOT NULL,
  `role`       ENUM ('user', 'admin') DEFAULT 'user' NOT NULL,
  `status`     TINYINT DEFAULT 1 NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NULL,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
  PRIMARY KEY (id),
  UNIQUE (email)
);

CREATE TABLE questions (
	`id` CHAR(36) NOT NULL,
  `owner_id` CHAR(36) NOT NULL,
	`content` VARCHAR(250) CHARACTER SET utf8mb4 NOT NULL,
	`category` CHAR(36) CHARACTER SET utf8mb4 NOT NULL,
	`type` ENUM ('multiple_choice', 'true_false') DEFAULT 'multiple_choice',
	`difficulty` ENUM ('easy', 'normal', 'hard') DEFAULT 'normal',
	`plus_score` INT NOT NULL,
  `minus_score` INT NOT NULL,
	`time` INT NOT NULL,
	`status` TINYINT DEFAULT 1,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);