DROP VIEW IF EXISTS v_best_selling_games;
DROP VIEW IF EXISTS v_total_revenue;

CREATE OR REPLACE VIEW v_best_selling_games AS
SELECT
    g.gameid,
    g.title,
    COUNT(o.orderid) AS total_sold,
    SUM(g.price) AS total_revenue
FROM games g
LEFT JOIN orders o ON g.gameid = o.gameid
GROUP BY g.gameid, g.title
ORDER BY total_sold DESC;

CREATE OR REPLACE VIEW v_total_revenue AS
SELECT
    SUM(amount) AS total_revenue,
    COUNT(paymentid) AS total_payments
FROM payments
WHERE status = 'PAID';