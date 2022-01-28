# SQLコーディング規約

- インデントは半角スペース4つ
- 予約語に対しては、大文字を使用
- テーブルエイリアスは母音を抜かした子音字を利用
  - ex.) idol -> idl

```sql
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
```