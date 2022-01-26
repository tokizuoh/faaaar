select
    *
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