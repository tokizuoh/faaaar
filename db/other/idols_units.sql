SELECT
    idl.id          AS id,
    idl.name        AS name,
    idl.age         AS age,
    idl.height      AS height,
    idl.birth_place AS birth_place,
    idl.blood_type  AS blood_type,
    unt.name        AS unit
FROM
    idol idl
INNER JOIN
    idol_unit idlunt
ON
    idl.id = idlunt.idol
INNER JOIN
    unit unt
ON
    idlunt.unit = unt.id;