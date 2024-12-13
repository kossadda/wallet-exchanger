CREATE TABLE IF NOT EXISTS users
(
  id               SERIAL         PRIMARY KEY         ,
  "username"       VARCHAR(255)   NOT NULL      UNIQUE,
  "password_hash"  VARCHAR(255)   NOT NULL            ,
  "email"          VARCHAR(255)   NOT NULL      UNIQUE
);
