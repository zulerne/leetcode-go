-- https://leetcode.com/problems/classes-with-at-least-5-students/description/

select class
from courses
group by class
having count(class) >= 5
