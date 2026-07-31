-- https://leetcode.com/problems/average-selling-price/description/

select
    p.product_id,
    coalesce(round(sum(u.units * p.price)::numeric / sum(u.units)::numeric, 2), 0) as average_price
from prices as p
left join unitssold as u
on u.product_id = p.product_id
    and u.purchase_date >= p.start_date
    and u.purchase_date <= p.end_date
group by p.product_id
