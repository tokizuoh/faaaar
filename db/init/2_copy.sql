copy idol(id,name,latin_alphabet,age,height,birth_place,birth_day,blood_type)
  from '/docker-entrypoint-initdb.d/idol_list.csv' with csv header;

copy unit(id,name)
  from '/docker-entrypoint-initdb.d/unit_list.csv' with csv header;

copy idol_unit(id,idol, unit)
  from '/docker-entrypoint-initdb.d/idol_unit.csv' with csv header;