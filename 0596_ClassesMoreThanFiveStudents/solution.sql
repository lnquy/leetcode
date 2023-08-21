SELECT class 
FROM classes
GROUP BY class 
	HAVING COUNT(DISTINCT student) >= 5;