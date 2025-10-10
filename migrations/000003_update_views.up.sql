-- Drop
DROP VIEW IF EXISTS v_best_selling_games;
DROP VIEW IF EXISTS v_total_revenue;

-- Game Terlaris
CREATE OR REPLACE VIEW v_best_selling_games AS
SELECT
    g.gameID,
    g.title AS nama_game,
    COUNT(o.orderID) AS total_terjual,
    CASE
        WHEN COUNT(o.orderID) > 0 THEN COUNT(o.orderID) * g.price
        ELSE 0
    END AS total_pendapatan
FROM games g
LEFT JOIN orders o ON g.gameID = o.gameID
GROUP BY g.gameID, g.title, g.price
ORDER BY total_terjual DESC;

-- View Total Pendapatan
CREATE OR REPLACE VIEW v_total_revenue AS
SELECT
    COALESCE(SUM(p.amount), 0) AS total_revenue,
    COALESCE(SUM(CASE WHEN LOWER(p.status) = 'unpaid' THEN p.amount ELSE 0 END), 0) AS outstanding_bills,
    COALESCE(SUM(CASE WHEN LOWER(p.status) = 'paid' THEN p.amount ELSE 0 END), 0) AS daily_income
FROM payments p;
