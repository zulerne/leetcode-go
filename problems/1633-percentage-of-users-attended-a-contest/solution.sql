-- https://leetcode.com/problems/percentage-of-users-attended-a-contest/description/

select
    r.contest_id,
    round(
       count(r.contest_id)::numeric
       / (select count(*) from users)
       * 100,
       2
    )
    as percentage
from register as r
join users as u
    on u.user_id = r.user_id
group by r.contest_id
order by percentage desc, r.contest_id
