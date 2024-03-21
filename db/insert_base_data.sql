-- Вставка записей в таблицу products
INSERT INTO products (id, name, price, stock_quantity)
VALUES (1, 'Ноутбук', 1000, 10),
       (2, 'Телевизор', 2000, 20),
       (3, 'Телефон', 500, 5),
       (4, 'Системный блок', 800, 8),
       (5, 'Часы', 300, 3),
       (6, 'Микрофон', 150, 2);

-- Вставка записей в таблицу shelves
INSERT INTO shelves (id, name)
VALUES (1, 'А'),
       (2, 'Б'),
       (3, 'В'),
       (4, 'З'),
       (5, 'Ж');

-- Вставка записей в таблицу orders
INSERT INTO orders (id, order_date)
VALUES (10, '2024-03-20 10:00:00'),
       (11, '2024-03-21 12:00:00'),
       (14, '2024-03-25 09:00:00'),
       (15, '2024-03-26 14:00:00');

-- Вставка записей в таблицу products_shelves
INSERT INTO products_shelves (product_id, shelves_id, is_primary)
VALUES (1, 1, true),
       (2, 1, true),
       (3, 2, true),
       (3, 3, false),
       (3, 4, false),
       (4, 5, true),
       (5, 5, true),
       (5, 1, false),
       (6, 5, true);

-- Вставка записей в таблицу order_items
INSERT INTO order_items (order_id, product_id, quantity)
VALUES (10, 1, 2),
       (10, 3, 1),
       (10, 6, 1),
       (11, 2, 3),
       (14, 1, 3),
       (14, 4, 4),
       (15, 5, 1);
