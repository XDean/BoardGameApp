CREATE TABLE IF NOT EXISTS `user_profiles`(  
  `user_id` BIGINT NOT NULL ,
  `nickname` VARCHAR(127) CHARACTER SET utf8,
  `gender` TINYINT(1),
  `avatar_url` TEXT,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) DEFAULT CHARSET=latin1;