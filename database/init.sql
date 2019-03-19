CREATE DATABASE IF NOT EXISTS gocrud_db;

CREATE TABLE IF NOT EXISTS gocrud_db.product (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255),
    PRIMARY KEY (id)
);
