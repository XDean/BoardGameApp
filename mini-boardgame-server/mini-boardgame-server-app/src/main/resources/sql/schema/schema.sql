CREATE TABLE IF NOT EXISTS `t_users`(
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(511) NOT NULL,
  `password` VARCHAR(511) NOT NULL,
  enabled BOOLEAN NOT NULL,
  UNIQUE INDEX `username` (`username`)
) DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `t_authorities` (
  `id` INT,
  `authority` VARCHAR(63) NOT NULL,
  UNIQUE INDEX ix_auth_id(`id`, `authority`),
  CONSTRAINT fk_authorities_users FOREIGN KEY(id) REFERENCES t_users(id)
) DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `t_user_profiles`(  
  `id` INT PRIMARY KEY,
  `nickname` VARCHAR(127) CHARACTER SET utf8,
  `male` TINYINT(1),
  `avatar_url` TEXT,
  CONSTRAINT `fk_profile_user_id` FOREIGN KEY (`id`) REFERENCES `t_users`(`id`)
) DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `t_game_rooms`(  
  `id` INT PRIMARY KEY,
  `game_name` VARCHAR(127),
  `room_name` VARCHAR(127),
  `player_count` INT,
  `created_time` TIMESTAMP,
  `board` TEXT
) DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `t_game_players`(  
  `id` INT PRIMARY KEY,
  `room_id` INT,
  `seat` INT DEFAULT -1,
  CONSTRAINT `fk_player_user_id` FOREIGN KEY (`id`) REFERENCES `t_users`(`id`),
  CONSTRAINT `fk_player_room_id` FOREIGN KEY (`room_id`) REFERENCES `t_game_rooms`(`id`)
) DEFAULT CHARSET=latin1;
