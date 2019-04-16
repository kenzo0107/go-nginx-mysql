-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id int NOT NULL PRIMARY KEY,
  name VARCHAR(20)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- +migrate Down
DROP TABLE IF EXISTS users;
