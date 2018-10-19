SELECT d.Name AS Department, e.Name AS Employee, e.Salary
FROM Department d, Employee e
WHERE d.Id = e.DepartmentId
  AND e.Salary = (SELECT MAX(Salary) FROM Employee e2 WHERE e2.DepartmentId = d.Id);

