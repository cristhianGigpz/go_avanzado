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

CREATE TABLE posts(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100),

	user_id INTEGER,

	CONSTRAINT fk_user
	FOREIGN KEY(user_id)
	REFERENCES users(id)
);

INSERT INTO posts(title,user_id)
VALUES('Mi primer post',1);

SELECT * FROM posts

CREATE TABLE accounts(
	id SERIAL PRIMARY KEY,
	balance FLOAT(24),
	user_id INTEGER,

	CONSTRAINT fk_user
	FOREIGN KEY(user_id)
	REFERENCES users(id)
)
SELECT * FROM accounts
INSERT INTO accounts (balance, user_id) VALUES (1500.12, 1)
INSERT INTO accounts (balance, user_id) VALUES (500.12, 2)

SELECT * FROM users
SELECT * FROM posts
SELECT * FROM roles
SELECT * FROM user_roles