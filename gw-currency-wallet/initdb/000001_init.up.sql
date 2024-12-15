CREATE TABLE IF NOT EXISTS users
(
  id               SERIAL         PRIMARY KEY         ,
  "username"       VARCHAR(255)   NOT NULL      UNIQUE,
  "password_hash"  VARCHAR(255)   NOT NULL            ,
  "email"          VARCHAR(255)   NOT NULL      UNIQUE
);

CREATE TABLE IF NOT EXISTS wallets
(
    id          SERIAL         PRIMARY KEY,
    user_id     INT            NOT NULL      UNIQUE,
    usd         NUMERIC(12, 2) NOT NULL      DEFAULT 0,
    rub         NUMERIC(12, 2) NOT NULL      DEFAULT 0,
    eur         NUMERIC(12, 2) NOT NULL      DEFAULT 0,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
