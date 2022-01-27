select
    i.id as id,
    i.name as name,
    i.age as age,
    i.height as height,
    i.birth_place as birth_place,
    i.blood_type as blood_type,
    u.name as unit
from
    idol i
inner join
    idol_unit iu
on
    i.id = iu.idol
inner join
    unit u
on
    iu.unit=u.id;