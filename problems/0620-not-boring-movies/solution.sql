-- https://leetcode.com/problems/not-boring-movies/description/

select *
from cinema as c
where c.id%2=1 and description != 'boring'
order by c.rating desc
