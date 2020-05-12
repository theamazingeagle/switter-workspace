--CREATE DATABASE switter;--
--USE switter;--
CREATE TABLE users(
    user_id SERIAL NOT NULL PRIMARY KEY,
    user_name VARCHAR(256),
    user_email VARCHAR(256),
    user_password VARCHAR(256)
);
CREATE TABLE messages(
    message_id SERIAL NOT NULL PRIMARY KEY,
    message_url VARCHAR(2048),
    message_userid BIGINT,
    message_date TIMESTAMP,
    --message_rootid VARCHAR(256),--
    message_text TEXT
);

ALTER TABLE messages ADD FOREIGN KEY (message_userid) REFERENCES users(user_id);
--alter table messages drop column message_rootid;--