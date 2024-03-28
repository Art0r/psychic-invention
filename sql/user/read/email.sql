-- based on model models.User

SELECT * FROM users 
WHERE email = $1 
OFFSET 0
LIMIT 1;