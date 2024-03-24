DROP TABLE IF EXISTS users;

-- based on model models.User

CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE
		);