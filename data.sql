CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100),
	email VARCHAR(100)
)

SELECT * FROM users
SELECT * FROM users WHERE id = 1
INSERT INTO users (id, name, email) values (1, 'cristhian joel', 'cjaceveodt@gmail.com');
UPDATE users SET name='Cristhian' WHERE id = 1
DELETE FROM users WHERE id = 1

SELECT * FROM users ORDER BY id DESC;
SELECT * FROM users LIMIT 1;
SELECT * FROM users WHERE email = 'gigpz@example.com'


SELECT * FROM users WHERE email LIKE '%.com%';

SELECT COUNT(*) FROM users;