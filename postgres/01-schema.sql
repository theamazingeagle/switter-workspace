--CREATE DATABASE switter;--
--USE switter;--
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(256),
    email VARCHAR(256),
    password VARCHAR(256),
    refresh_token VARCHAR(256)
);

CREATE TABLE messages(
    id SERIAL NOT NULL PRIMARY KEY,
    msg TEXT,
    user_id BIGINT,
    msg_date TIMESTAMP
);

ALTER TABLE users ALTER COLUMN refresh_token SET DEFAULT '';
ALTER TABLE messages ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE messages ALTER COLUMN msg_date SET DEFAULT Now();