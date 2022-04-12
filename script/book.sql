CREATE SCHEMA IF NOT EXISTS books;

USE books;

CREATE TABLE books (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL,
	UNIQUE (title)
)