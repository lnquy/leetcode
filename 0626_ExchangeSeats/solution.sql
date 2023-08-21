SELECT id,
CASE 
    WHEN id%2 = 0 THEN(SELECT student FROM seat WHERE id = (s.id-1) )  
    WHEN id%2 != 0 AND id < (SELECT count(student) FROM seat) THEN (SELECT student FROM seat WHERE id = (s.id+1) )
    ELSE student
END AS student
FROM seat s;
