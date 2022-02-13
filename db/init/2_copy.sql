USE shiny_colors_db;

LOAD DATA INFILE
  '/docker-entrypoint-initdb.d/idol_list.csv'
INTO TABLE
  idol
FIELDS
TERMINATED BY ','
IGNORE 1 ROWS
SET character_set_database=utf-8;

LOAD DATA INFILE
  '/docker-entrypoint-initdb.d/unit_list.csv'
INTO TABLE
  unit
FIELDS
TERMINATED BY ','
IGNORE 1 ROWS;

LOAD DATA INFILE
  '/docker-entrypoint-initdb.d/idol_unit.csv'
INTO TABLE
  idol_unit
FIELDS
TERMINATED BY ','
IGNORE 1 ROWS;