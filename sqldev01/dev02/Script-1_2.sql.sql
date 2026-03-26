DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
	id INT UNSIGNED  NOT NULL PRIMARY KEY,
	user_id INTEGER NULL,
	products_count INTEGER NULL,
	sum INTEGER NULL,
	status VARCHAR(20) NULL
);

INSERT INTO orders (id, user_id, products_count, sum, status)
VALUES
	(1, 1, 2, 1300, 'new'),
    (2, 18, 1, 10000, 'cancelled'),
    (3, 11, 1, 2140, 'in_progress'),
    (4, 145, 5, 6800, 'new'),
    (5, 23, 1, 999, 'new'),
    (6, 1, 2, 7690, 'cancelled'),
    (7, 17, 1, 1600, 'new'),
    (8, 5, 4, 400, 'delivery'),
    (9, 2355, 1, 1450, 'new'),
    (10, 13, 7, 13000, 'cancelled'),
    (11, 144, 6, 3000, 'cancelled');

SELECT * FROM orders WHERE status != 'cancelled';

SELECT id, sum FROM orders WHERE products_count > 3;

SELECT * FROM orders WHERE status = 'cancelled';

SELECT * FROM orders WHERE status IN ('cancelled', 'returned');

SELECT * FROM orders WHERE sum > 3000 OR products_count >= 3;

SELECT * FROM orders WHERE sum >= 3000 AND products_count < 3;

SELECT * FROM orders
WHERE status = 'cancelled'
	AND sum BETWEEN 3000 AND 10000;

SELECT * FROM orders
WHERE status = 'cancelled'
	AND sum NOT BETWEEN 3000 AND 10000;

SELECT * FROM orders
WHERE status = 'cancelled'
AND (sum < 3000 OR sum > 10000);
