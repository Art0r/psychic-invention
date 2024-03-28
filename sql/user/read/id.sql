-- based on model models.User

SELECT * FROM users 
WHERE id = $1 
OFFSET 0
LIMIT 1;