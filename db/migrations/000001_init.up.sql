CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL         PRIMARY KEY,
    username        VARCHAR(255)   NOT NULL UNIQUE,
    password_hash   VARCHAR(255)   NOT NULL,
    email           VARCHAR(255)   NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS wallets
(
    id              SERIAL         PRIMARY KEY,
    user_id         INT            NOT NULL UNIQUE,
    usd             NUMERIC(12, 2) NOT NULL DEFAULT 0,
    rub             NUMERIC(12, 2) NOT NULL DEFAULT 0,
    eur             NUMERIC(12, 2) NOT NULL DEFAULT 0,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS currency
(
    id              SERIAL         PRIMARY KEY,
    output          VARCHAR(15)    NOT NULL UNIQUE,
    usd             NUMERIC(12, 7) NOT NULL DEFAULT 0,
    rub             NUMERIC(12, 7) NOT NULL DEFAULT 0,
    eur             NUMERIC(12, 7) NOT NULL DEFAULT 0
);

INSERT INTO currency (output, usd, rub, eur)
VALUES
    ('usd', 1, 0.0097, 1.05),
    ('rub', 103.6, 1, 108.89),
    ('eur', 0.95, 0.0092, 1);
