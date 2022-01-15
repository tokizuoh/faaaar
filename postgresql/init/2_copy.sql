copy idol(id,name,age,height,birth_place,birth_day,blood_type,unit)
  from '/docker-entrypoint-initdb.d/idol_list.csv' with csv header;