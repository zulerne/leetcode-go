-- https://leetcode.com/problems/rising-temperature/description/

select today.id
from weather as today
join weather as yesterday
on today.recordDate - interval '1 day' = yesterday.recordDate
where today.temperature > yesterday.temperature
