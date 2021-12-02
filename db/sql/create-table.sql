DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
    `user_id` INT(20) AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL,
    `email` VARCHAR(40) NOT NULL,
    `password` VARCHAR(80),
    PRIMARY KEY (`user_id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

DROP TABLE IF EXISTS `todos`;

CREATE TABLE IF NOT EXISTS `todos` (
    `todo_id` INT(40) AUTO_INCREMENT,
    `title` VARCHAR(20) NOT NULL,
    `content` VARCHAR(20) NOT NULL,
    `done` BOOLEAN NOT NULL,
    `user_id` INT(20) NOT NULL,
    PRIMARY KEY (`todo_id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
