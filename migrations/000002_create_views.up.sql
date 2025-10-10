-- View History Pembelian Customer
CREATE OR REPLACE VIEW v_customer_purchase_history AS
SELECT
    o.OrderID,
    c.CustomerID,
    c.Name AS customer_name,
    c.Email AS customer_email,
    g.GameID,
    g.Title AS game_title,
    g.Price,
    o.Status AS payment_status,
    o.CreatedAt AS order_date
FROM Orders o
JOIN Customers c ON o.CustomerID = c.CustomerID
JOIN Games g ON o.GameID = g.GameID
WHERE o.Status = 'PAID'
ORDER BY o.OrderID;

-- View Game Terlaris
CREATE OR REPLACE VIEW v_best_selling_games AS
SELECT
    g.gameID,
    g.title,
    COUNT(o.orderID) AS total_sold,
    SUM(g.price) AS total_revenue
FROM games g
LEFT JOIN orders o ON g.gameID = o.gameID
GROUP BY g.gameID, g.title
ORDER BY total_sold DESC;

-- View Total Pendapatan
CREATE OR REPLACE VIEW v_total_revenue AS
SELECT
    SUM(Amount) AS total_revenue,
    COUNT(PaymentID) AS total_payments
FROM payments
WHERE Status = 'PAID';

-- View Summary
CREATE OR REPLACE VIEW v_summary AS
SELECT
    (SELECT COUNT(*) FROM customers) AS total_customers,
    (SELECT COUNT(*) FROM games) AS total_games,
    (SELECT COUNT(*) FROM orders) AS total_orders,
    (SELECT COUNT(*) FROM payments) AS total_payments,
    (SELECT SUM(amount) FROM payments WHERE status = 'PAID') AS total_revenue;