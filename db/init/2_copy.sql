USE shiny_colors_db;

LOAD DATA INFILE
  '/docker-entrypoint-initdb.d/idol_list.csv'
INTO TABLE
  idol
FIELDS
TERMINATED BY ','
IGNORE 1 ROWS;  /* 追加 */

copy idol_unit(id,idol, unit)
  from '/docker-entrypoint-initdb.d/idol_unit.csv' with csv header;