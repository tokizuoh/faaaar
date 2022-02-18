CREATE DATABASE shiny_colors_db;
USE shiny_colors_db;

/* from idol_list.csv */
CREATE TABLE idol
(
    `id` INT(2) AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` VARCHAR(30) NOT NULL,
    `latin_alphabet` VARCHAR(30) NOT NULL,
    `age` INT(2) NOT NULL,
    `height` INT(3) NOT NULL,
    `birth_place` VARCHAR(30) NOT NULL,
    `birth_day` VARCHAR(5) NOT NULL,
    `blood_type` VARCHAR(2) NOT NULL
);

/* from unit_list.csv */
CREATE TABLE unit
(
    `id` VARCHAR(10) NOT NULL PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL
);

/* from idol_unit.csv */
CREATE TABLE idol_unit
(
    `id` INT(3) AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `idol` VARCHAR(30) NOT NULL,
    `unit` VARCHAR(30) NOT NULL
);