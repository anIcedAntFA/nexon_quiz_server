CREATE TABLE users (
    `id`         CHAR(36) NOT NULL,
    `email`      VARCHAR(50) NOT NULL,
    `password`   VARCHAR(50) NOT NULL,
    `salt`       VARCHAR(50) NULL,
    `user_name`  VARCHAR(50) NOT NULL,
    `role`       ENUM ('user', 'admin') DEFAULT 'user' NOT NULL,
    `status`     TINYINT DEFAULT 1 NOT NULL,
    `createdAt`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP                             NULL,
    `updatedAt`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE (email)
);