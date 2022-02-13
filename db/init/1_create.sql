CREATE DATABASE shiny_colors_db;
USE shiny_colors_db;

/* from idol_list.csv */
CREATE TABLE idol
(
    `id` INT(2) NOT NULL PRIMARY KEY,
    `name` VARCHAR(30),
    `latin_alphabet` VARCHAR(30),
    `age` INT(2),
    `height` INT(3),
    `birth_place` VARCHAR(30),
    `birth_day` VARCHAR(5),
    `blood_type` VARCHAR(2)
);

/* from unit_list.csv */
CREATE TABLE unit
(
    `id` VARCHAR(10) NOT NULL PRIMARY KEY,
    `name` VARCHAR(30)
);

/* from idol_unit.csv */
CREATE TABLE idol_unit
(
    `id` INT(3) NOT NULL PRIMARY KEY,
    `idol` VARCHAR(30),
    `unit` VARCHAR(30)
);