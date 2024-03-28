-- based on model models.User

SELECT (id, name, email) FROM users
ORDER BY name DESC
OFFSET 0
LIMIT 100;