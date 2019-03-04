CREATE TABLE IF NOT EXISTS users(
  id int PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(511) NOT NULL,
  `password` VARCHAR(511) NOT NULL,
  enabled BOOLEAN NOT NULL,
  UNIQUE INDEX `username` (`username`)
) DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS authorities (
      username VARCHAR(511) NOT NULL,
      authority VARCHAR(63) NOT NULL,
      UNIQUE INDEX ix_auth_username(`username`, `authority`),
      CONSTRAINT fk_authorities_users FOREIGN KEY(username) REFERENCES users(username)
) DEFAULT CHARSET=latin1;

/*
CREATE TABLE IF NOT EXISTS groups (
  id int PRIMARY KEY AUTO_INCREMENT,
  group_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS group_authorities (
  group_id int NOT NULL,
  authority VARCHAR(50) NOT NULL,
  CONSTRAINT fk_group_authorities_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS group_members (
  id int PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(512) NOT NULL,
  group_id BIGINT NOT NULL,
  CONSTRAINT fk_group_members_users FOREIGN KEY(username) REFERENCES users(username),
  CONSTRAINT fk_group_members_group FOREIGN KEY(group_id) REFERENCES groups(id)
);
*/