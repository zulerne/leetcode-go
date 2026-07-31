-- https://leetcode.com/problems/average-time-of-process-per-machine/description/

select
    e.machine_id,
    round(avg(e.timestamp-s.timestamp)::numeric, 3) as processing_time
from activity as e
join activity as s
    on e.activity_type = 'end' and s.activity_type = 'start'
        and e.machine_id = s.machine_id
        and e.process_id = s.process_id
group by e.machine_id
