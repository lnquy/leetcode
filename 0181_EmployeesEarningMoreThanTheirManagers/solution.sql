SELECT Name as Employee
FROM Employee AS e
WHERE Salary > (SELECT Salary FROM Employee WHERE Id = e.ManagerId);
