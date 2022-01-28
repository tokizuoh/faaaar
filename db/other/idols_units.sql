SELECT
    i.id AS id,
    i.name AS name,
    i.age AS age,
    i.height AS height,
    i.birth_place AS birth_place,
    i.blood_type AS blood_type,
    u.name AS unit
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