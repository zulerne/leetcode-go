-- https://leetcode.com/problems/user-activity-for-the-past-30-days-i/description/

select a.activity_date as day, (count(distinct a.user_id)) as active_users
from activity as a
where a.activity_date between
    '2019-07-27'::date - interval '30 days' and
    '2019-07-27'::date
group by a.activity_date
