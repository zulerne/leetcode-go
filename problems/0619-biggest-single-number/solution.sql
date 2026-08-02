-- https://leetcode.com/problems/biggest-single-number/description/

select max(num) as num
from (
    select num
    from mynumbers
    group by num
    having count(*) = 1
)
