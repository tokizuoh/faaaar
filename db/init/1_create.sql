create table idol
(
    id integer not null,
    name varchar not null,
    age integer not null,
    height integer not null,
    birth_place varchar not null,
    birth_day varchar not null,
    blood_type varchar not null,
    unit varchar,
    PRIMARY KEY (id)
);
comment on table idol is 'idol_list.csv';

create table unit
(
    id varchar not null,
    name varchar not null,
    PRIMARY KEY (id)
);
comment on table idol is 'unit_list.csv';