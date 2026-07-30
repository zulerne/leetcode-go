-- https://leetcode.com/problems/product-sales-analysis-i/description/

select p.product_name, s.year, s.price
from Product as p
join Sales as s
    on s.product_id = p.product_id;
