DROP TABLE IF EXISTS products;
CREATE TABLE products (
	id INT UNSIGNED NOT NULL,
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
    (6, 'Телевизор', 7, 31740);

SELECT * FROM products;

SELECT name, price FROM products;

SELECT * FROM products WHERE price < 3000

INSERT INTO products (id, name, count, price)
VALUES
	(7, 'Тостер', 2, 2500),
	(8, 'Принтер', 4, 3000);

SELECT name, price FROM products WHERE price >= 10000; 

SELECT name FROM products WHERE count = 0;

SELECT name, price FROM products WHERE price <= 4000;
