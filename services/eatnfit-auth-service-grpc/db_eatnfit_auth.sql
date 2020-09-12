# Use this file for db

CREATE DATABASE db_eatnfit_auth;
USE db_eatnfit_auth;

CREATE TABLE tb_user (
     user_id VARCHAR(36) PRIMARY KEY NOT NULL,
     user_email VARCHAR(100) NOT NULL,
     user_password VARCHAR(255) NOT NULL,
     user_f_name VARCHAR(100) NOT NULL,
     user_l_name VARCHAR(100) NOT NULL,
     user_gender INT NOT NULL,
     user_photo VARCHAR(255) NOT NULL,
     user_balance INT NULL DEFAULT 0,
     user_level INT NOT NULL,
     user_status INT NOT NULL
);

CREATE TABLE tb_level (
    level_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    level_name VARCHAR(50) NOT NULL,
    level_status INT NOT NULL
);

INSERT INTO tb_level VALUES (NULL, 'Admin', 1),
                            (NULL, 'User', 1),
                            (NULL, 'Driver', 1);
