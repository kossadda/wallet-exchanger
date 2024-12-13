CREATE TABLE IF NOT EXISTS users
(
  id               SERIAL         PRIMARY KEY         ,
  "name"           VARCHAR(255)   NOT NULL            ,
  "username"       VARCHAR(255)   NOT NULL      UNIQUE,
  "password_hash"  VARCHAR(255)   NOT NULL
);
