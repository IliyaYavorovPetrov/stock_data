CREATE TABLE stocks.users
(
    user_uuid VARCHAR(36) NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);