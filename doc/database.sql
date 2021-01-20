CREATE DATABASE demo;

USE demo;

CREATE TABLE users(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(60) NOT NULL,
    due_date varchar(10) NOT NULL,
    complt_date varchar(10) NOT NULL,
    status tinyint(2) NOT NULL,
    PRIMARY KEY(id)
);

