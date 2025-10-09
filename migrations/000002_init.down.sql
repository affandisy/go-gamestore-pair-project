DELETE FROM customers WHERE email = 'admin@gmail.com';

ALTER TABLE customers DROP COLUMN role;

ALTER TABLE customers ALTER COLUMN password TYPE VARCHAR(255);
