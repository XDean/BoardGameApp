CREATE TABLE IF NOT EXISTS `users_profiles`(  
  `username` VARCHAR(512) NOT NULL,
  `nickname` VARCHAR(128),
  `gender` TINYINT(1),
  `avatar_url` TEXT,
  PRIMARY KEY (`username`),
  CONSTRAINT `fk_username` FOREIGN KEY (`username`) REFERENCES `users`(`username`)
);