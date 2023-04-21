CREATE TABLE users (
  `id`             CHAR(36) NOT NULL,
  `email`          VARCHAR(50) NOT NULL,
  `password`       VARCHAR(50) NOT NULL,
  `salt`           VARCHAR(50) NULL,
  `username`       VARCHAR(50) NOT NULL,
  `role`           ENUM ('user', 'admin') DEFAULT 'user' NOT NULL,
  `is_deleted`     TINYINT DEFAULT 0 NOT NULL,
  `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NULL,
  `updated_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE (email)
);

CREATE TABLE questions (
	`id` CHAR(36) NOT NULL,
  `owner_id` CHAR(36) NOT NULL,
	`content` VARCHAR(250) CHARACTER SET utf8mb4 NOT NULL,
	`category` CHAR(36) CHARACTER SET utf8mb4 NOT NULL,
	`type` ENUM ("single_choice", "multiple_choice", "true_false") DEFAULT 'single_choice',
	`difficulty` ENUM ('easy', 'normal', 'hard') DEFAULT 'normal',
	`plus_score` INT NOT NULL,
  `minus_score` INT NOT NULL,
	`time` INT NOT NULL,
	`is_deleted` TINYINT DEFAULT 0,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

CREATE TABLE answers (
	`id` CHAR(36) NOT NULL,
  `question_id` CHAR(36) NOT NULL,
	`content` VARCHAR(250) CHARACTER SET utf8mb4 NOT NULL,
  `correct` TINYINT NOT NULL,
	`is_deleted` TINYINT DEFAULT 0,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);