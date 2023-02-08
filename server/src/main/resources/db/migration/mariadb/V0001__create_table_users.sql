CREATE DATABASE IF NOT EXISTS stocks;

USE stocks;

CREATE TABLE
    IF NOT EXISTS stocks.persons
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    first_name  VARCHAR(255),
    last_name   VARCHAR(255),
    national_id VARCHAR(255) UNIQUE
);