DROP TABLE IF EXISTS products;

CREATE TABLE products (
	id INT UNSIGNED NOT NULL PRIMARY KEY,
	name VARCHAR(255) NULL,
	count INTEGER NULL,
	price INTEGER NULL
);

INSERT INTO products (id, name, count, price)
VALUES 
	(1, 'Стиральная машина', 5, 10000),
	(2, 'Холодильник', 0, 10000),
    (3, 'Микроволновка', 3, 4000),
    (4, 'Пылесос', 2, 4500),
    (5, 'Вентилятор', 0, 700),
    (6, 'Телевизор', 7, 31740),
    (7, 'Тостер', 2, 2500),
    (8, 'Принтер', 4, 3000);

SELECT * FROM products;

SELECT * FROM products ORDER BY price;

SELECT name, price FROM products ORDER BY price DESC;

SELECT * FROM products WHERE price >= 5000 ORDER BY price DESC;

SELECT name, count, price FROM products WHERE price < 3000 ORDER BY name;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id INT UNSIGNED NOT NULL PRIMARY KEY,
	first_name VARCHAR(50) NULL,
	last_name VARCHAR(50) NULL,
	birthday DATE NULL
);

INSERT INTO users (id, first_name, last_name, birthday)
VALUES 
	(1, 'Дмитрий', 'Петров', '2000-03-14'),
    (2, 'Ольга', 'Антонова', '1999-12-01'),
    (3, 'Сергей', 'Васильев', '2002-02-20'),
    (4, 'Константин', 'Степаниденко', '2004-03-07'),
    (5, 'Алена', 'Шикова', '1999-08-17'),
    (6, 'Василина', 'Парамонова', '2000-10-10'),
    (7, 'Александр', 'Пузаков', '2002-02-20'),
    (8, 'Алина', 'Антонова', '2002-01-01');

SELECT last_name, first_name FROM users ORDER BY last_name, first_name;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT UNSIGNED NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NULL,
    last_name VARCHAR(50) NULL,
    birthday DATE NULL,
    salary INTEGER NULL,
    job VARCHAR(50) NULL
);

INSERT INTO users (id, first_name, last_name, birthday, salary, job)
VALUES
    (1, 'Дмитрий', 'Петров', '2000-03-14', 25000, 'офис-менеджер'),
    (2, 'Ольга', 'Антонова', '1999-12-01', 41000, 'дизайнер'),
    (3, 'Сергей', 'Васильев', '2002-02-20', 40000, 'младший программист'),
    (4, 'Констанин', 'Степаниденко', '2004-03-07', 30000, 'водитель'),
    (5, 'Алена', 'Шикова', '1999-08-17', 53000, 'фотограф'),
    (6, 'Василина', 'Парамнова', '2000-10-10', 28000, 'секретарь'),
    (7, 'Александр', 'Пузаков', '2002-02-20', 120000, 'ведущий программист'),
    (8, 'Алина', 'Антонова', '2002-01-01', 40000, 'верстальщик');

SELECT * FROM users WHERE salary >= 40000 ORDER BY salary DESC, first_name;

SELECT * FROM users WHERE salary > 0 AND salary < 30000 ORDER BY birthday;
