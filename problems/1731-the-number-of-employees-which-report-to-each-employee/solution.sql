-- https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/description/

select
    m.employee_id,
    m.name,
    count(*) as reports_count,
    round(avg(e.age)) as average_age
from employees as m
join employees as e
    on e.reports_to = m.employee_id
group by m.employee_id, m.name
order by m.employee_id
