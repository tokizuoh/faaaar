SELECT
    unt.name        AS unit_name,
    idl.name          AS idol_name,
    idl.id AS idol_id
    
FROM
    idol idl
INNER JOIN
    idol_unit idlunt
ON
    idl.id = idlunt.idol
INNER JOIN
    unit unt
ON
    idlunt.unit = unt.id