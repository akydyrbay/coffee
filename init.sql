-- 1. Create Type Statements

-- ENUM for order status
CREATE TYPE order_status AS ENUM (
    'pending',      -- Order is awaiting preparation
    'in_progress',  -- Order is being prepared
    'completed',    -- Order is finished
    'canceled'      -- Order is canceled
);

-- ENUM for unit of measurement (UOM)
CREATE TYPE unit_of_measurement AS ENUM (
    'kg', 'g', 'lb', 'oz', 'liter', 'ml', 'piece'
);

-- Composite type for menu item ingredients
CREATE TYPE menu_item_ingredients AS (
    ingredient_id INT,
    quantity DECIMAL,
    uom unit_of_measurement
);

-- 2. Create Table Statements

-- Core Tables

-- Table for menu items (e.g., coffee, snacks)
CREATE TABLE menu_items (
    menu_item_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(50),
    is_available BOOLEAN DEFAULT TRUE
);

-- Table for inventory items (e.g., coffee beans, milk)
CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    uom unit_of_measurement NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

-- Table for orders
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_name VARCHAR(100),
    order_status order_status NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL
);

-- Junction Table for menu item ingredients (many-to-many relationship)
CREATE TABLE menu_item_ingredients_junction (
    menu_item_id INT REFERENCES menu_items(menu_item_id) ON DELETE CASCADE,
    inventory_id INT REFERENCES inventory(inventory_id) ON DELETE CASCADE,
    quantity DECIMAL NOT NULL,
    uom unit_of_measurement NOT NULL,
    PRIMARY KEY (menu_item_id, inventory_id)
);

-- History Tables

-- Table to track order status changes (order status history)
CREATE TABLE order_status_history (
    history_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(order_id) ON DELETE CASCADE,
    status order_status NOT NULL,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for tracking price history (for menu items)
CREATE TABLE price_history (
    history_id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(menu_item_id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL,
    effective_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for inventory transactions (inventory movements like restocks, usage)
CREATE TABLE inventory_transactions (
    transaction_id SERIAL PRIMARY KEY,
    inventory_id INT REFERENCES inventory(inventory_id) ON DELETE CASCADE,
    quantity DECIMAL(10, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,  -- 'restock', 'usage'
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Create Index Statements

-- Index on menu items (for quick lookups by category)
CREATE INDEX idx_menu_items_category ON menu_items (category);

-- Index on orders (for quick search by order status)
CREATE INDEX idx_orders_status ON orders (order_status);

-- Full text search index for menu item names and descriptions
CREATE INDEX idx_menu_items_full_text ON menu_items USING gin(to_tsvector('english', name || ' ' || description));

-- Composite index for inventory transactions based on inventory_id and transaction type
CREATE INDEX idx_inventory_transactions_type ON inventory_transactions (inventory_id, transaction_type);

-- 4. Mock Data Inserts

-- Inserting sample menu items
INSERT INTO menu_items (name, description, price, category)
VALUES 
    ('Espresso', 'Strong black coffee', 3.50, 'coffee'),
    ('Americano', 'Espresso with hot water', 4.00, 'coffee'),
    ('Cappuccino', 'Espresso with steamed milk and foam', 4.50, 'coffee'),
    ('Latte', 'Espresso with steamed milk', 4.00, 'coffee'),
    ('Mocha', 'Espresso with chocolate and steamed milk', 4.75, 'coffee'),
    ('Cinnamon Roll', 'Freshly baked cinnamon roll', 2.50, 'pastry'),
    ('Bagel', 'Toasted bagel with cream cheese', 3.00, 'pastry'),
    ('Croissant', 'Flaky buttery croissant', 2.75, 'pastry'),
    ('Iced Coffee', 'Cold brewed coffee', 3.25, 'coffee'),
    ('Iced Latte', 'Cold latte', 4.25, 'coffee');

-- Inserting sample inventory items
INSERT INTO inventory (name, quantity, uom, price)
VALUES 
    ('Coffee Beans', 50, 'kg', 5.00),
    ('Milk', 100, 'liter', 2.00),
    ('Sugar', 10, 'kg', 1.50),
    ('Flour', 20, 'kg', 0.80),
    ('Cinnamon', 2, 'kg', 8.00),
    ('Chocolate Syrup', 5, 'liter', 6.00),
    ('Butter', 15, 'kg', 3.00),
    ('Eggs', 200, 'piece', 0.25),
    ('Cream Cheese', 10, 'kg', 4.00),
    ('Vanilla Syrup', 3, 'liter', 7.00);

-- Inserting sample orders
INSERT INTO orders (customer_name, order_status, total_amount)
VALUES 
    ('John Doe', 'pending', 15.50),
    ('Jane Smith', 'completed', 10.75),
    ('Jim Brown', 'in_progress', 12.00),
    ('Emily Clark', 'canceled', 0.00),
    ('Robert White', 'pending', 22.00),
    ('Alice Green', 'completed', 19.50),
    ('Michael King', 'in_progress', 18.25),
    ('Sarah Blue', 'completed', 20.00),
    ('David Black', 'pending', 11.00),
    ('Mary Pink', 'in_progress', 17.50);

-- Inserting sample order status history
INSERT INTO order_status_history (order_id, status)
VALUES 
    (1, 'pending'),
    (2, 'completed'),
    (3, 'in_progress'),
    (4, 'canceled'),
    (5, 'pending'),
    (6, 'completed'),
    (7, 'in_progress'),
    (8, 'completed'),
    (9, 'pending'),
    (10, 'in_progress');

-- Inserting sample price history for menu items
INSERT INTO price_history (menu_item_id, price)
VALUES 
    (1, 3.50),
    (2, 4.00),
    (3, 4.50),
    (4, 4.00),
    (5, 4.75),
    (6, 2.50),
    (7, 3.00),
    (8, 2.75),
    (9, 3.25),
    (10, 4.25);

-- Inserting sample inventory transactions (restocks and usage)
INSERT INTO inventory_transactions (inventory_id, quantity, transaction_type)
VALUES 
    (1, 10, 'restock'),
    (2, 20, 'usage'),
    (3, 5, 'usage'),
    (4, 10, 'restock'),
    (5, 2, 'restock'),
    (6, 3, 'usage'),
    (7, 7, 'restock'),
    (8, 50, 'usage'),
    (9, 5, 'restock'),
    (10, 2, 'usage');

-- 5. Testing Coverage

-- Test Full text search on menu items
-- Example query: Search for "latte"
SELECT * FROM menu_items WHERE to_tsvector('english', name || ' ' || description) @@ to_tsquery('english', 'latte');

-- Test date range query for orders (e.g., orders within the last 7 days)
SELECT * FROM orders WHERE order_date >= CURRENT_DATE - INTERVAL '7 days';

-- Test status transition in order_status_history (e.g., order status change for order 1)
SELECT * FROM order_status_history WHERE order_id = 1;

-- Test inventory calculations (e.g., total inventory quantity for "Milk")
SELECT SUM(quantity) FROM inventory WHERE name = 'Milk';

-- Test price history tracking (e.g., price history for item 1)
SELECT * FROM price_history WHERE menu_item_id = 1;
