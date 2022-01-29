SELECT
    idl.id          AS id,
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