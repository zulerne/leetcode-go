-- https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/description/

select EmployeeUNI.unique_id, Employees.name
from EmployeeUNI
right join Employees
    on Employees.id = EmployeeUNI.id
