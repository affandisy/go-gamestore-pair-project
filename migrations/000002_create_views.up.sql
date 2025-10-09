-- View History Pembelian Customer
CREATE OR REPLACE VIEW v_customer_purchase_history AS
SELECT
    o.OrderID,
    c.CustomerID,
    c.name AS customer_name,
    c.email AS customer_email,
    g.GameID,
    g.title AS game_title,
    g.price,
    COALESCE(p.status, 'UNPAID') AS payment_status,
    p.amount AS paid_amount,
    o.CreatedAt AS order_date
FROM orders o
JOIN customers c ON o.customerID = c.customerID
JOIN games g ON o.gameID = g.gameID
LEFT JOIN payments p ON o.orderID = p.orderID
ORDER BY o.orderID;

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
WHERE Status == "PAID"

-- View Summary
CREATE OR REPLACE VIEW v_summary AS
SELECT
    (SELECT COUNT(*) FROM customers) AS total_customers,
    (SELECT COUNT(*) FROM games) AS total_games,
    (SELECT COUNT(*) FROM orders) AS total_orders,
    (SELECT COUNT(*) FROM payments) AS total_payments,
    (SELECT SUM(amount) FROM payments WHERE status == "PAID") AS total_revenue;