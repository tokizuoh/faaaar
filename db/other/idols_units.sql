SELECT
    i.id as id,
    i.name as name,
    i.age as age,
    i.height as height,
    i.birth_place as birth_place,
    i.blood_type as blood_type,
    u.name as unit
FROM
    idol i
INNER JOIN
    idol_unit iu
ON
    i.id = iu.idol
INNER JOIN
    unit u
ON
    iu.unit=u.id;