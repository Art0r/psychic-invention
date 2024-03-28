-- based on model models.User

SELECT * FROM users 
WHERE name = $1 
OFFSET 0
LIMIT 1;