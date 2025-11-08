INSERT INTO orders (first_name, last_name, email, created_at, updated_at) VALUES
('John', 'Doe', 'john.doe@example.com', '2024-01-15 10:30:00', '2024-01-15 10:30:00'),
('Jane', 'Smith', 'jane.smith@example.com', '2024-01-16 14:20:00', '2024-01-16 14:20:00'),
('Mike', 'Johnson', 'mike.johnson@example.com', '2024-01-17 09:15:00', '2024-01-17 09:15:00'),
('Sarah', 'Wilson', 'sarah.wilson@example.com', '2024-01-18 16:45:00', '2024-01-18 16:45:00'),
('David', 'Brown', 'david.brown@example.com', '2024-01-19 11:20:00', '2024-01-19 11:20:00');

INSERT INTO order_items (order_id, product_title, price, quantity) VALUES
-- Order 1: John Doe's order
(1, 'Wireless Headphones', 199.99, 1),
(1, 'Phone Case', 25.50, 2),
(1, 'Screen Protector', 15.00, 1),

-- Order 2: Jane Smith's order
(2, 'Laptop', 1299.99, 1),
(2, 'Laptop Bag', 59.99, 1),

-- Order 3: Mike Johnson's order
(3, 'Smart Watch', 299.99, 1),
(3, 'Watch Band', 35.00, 2),

-- Order 4: Sarah Wilson's order
(4, 'Tablet', 499.99, 1),
(4, 'Tablet Cover', 29.99, 1),
(4, 'Stylus Pen', 45.00, 1),

-- Order 5: David Brown's order
(5, 'Gaming Mouse', 79.99, 1),
(5, 'Mechanical Keyboard', 149.99, 1),
(5, 'Mouse Pad', 19.99, 1);