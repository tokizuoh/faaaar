CREATE DATABASE shiny_colors_db;
USE shiny_colors_db;

/* from idol_list.csv */
CREATE TABLE idol
(
    id INT(2),
    name VARCHAR(30),
    latin_alphabet VARCHAR(30),
    age INT(2),
    height INT(3),
    birth_place VARCHAR(30),
    birth_day VARCHAR(5),
    blood_type VARCHAR(2),
    PRIMARY KEY (id)
);

create table unit
(
    id varchar not null,
    name varchar not null,
    PRIMARY KEY (id)
);
comment on table idol is 'unit_list.csv';

create table idol_unit
(
    id integer not null,
    idol integer not null,
    unit varchar not null,
    PRIMARY KEY (id)
);
comment on table idol is 'idol_unit.csv';