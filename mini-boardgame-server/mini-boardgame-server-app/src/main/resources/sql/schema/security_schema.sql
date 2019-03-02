CREATE TABLE IF NOT EXISTS users(
      `username` VARCHAR(512) NOT NULL PRIMARY KEY,
      `password` VARCHAR(512) NOT NULL,
      enabled BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS authorities (
      username VARCHAR(512) NOT NULL,
      authority VARCHAR(50) NOT NULL,
      UNIQUE INDEX ix_auth_username(`username`, `authority`),
      CONSTRAINT fk_authorities_users FOREIGN KEY(username) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS groups (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  group_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS group_authorities (
  group_id BIGINT NOT NULL,
  authority VARCHAR(50) NOT NULL,
  CONSTRAINT fk_group_authorities_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS group_members (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(512) NOT NULL,
  group_id BIGINT NOT NULL,
  CONSTRAINT fk_group_members_users FOREIGN KEY(username) REFERENCES users(username),
  CONSTRAINT fk_group_members_group FOREIGN KEY(group_id) REFERENCES groups(id)
);