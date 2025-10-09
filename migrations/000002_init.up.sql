ALTER TABLE customers ADD COLUMN role VARCHAR(255) DEFAULT 'customer';
INSERT INTO customers (name, email, password, role) VALUES ('admin', 'admin@gmail.com', 'admin123', 'admin');

ALTER TABLE customers ALTER COLUMN password TYPE VARCHAR(255);