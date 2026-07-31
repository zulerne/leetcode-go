-- https://leetcode.com/problems/queries-quality-and-percentage/description/

select
    q.query_name,
    round(avg(q.rating::numeric / q.position::numeric), 2)
        as quality,
    round(count(*) filter (where q.rating <3)::numeric / count(*) * 100, 2)
        as poor_query_percentage
from queries as q
group by q.query_name
