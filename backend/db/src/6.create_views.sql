-- Widok zmaterializowany: hotel_performance_summary
-- Agreguje metryki wydajności hoteli: rating, liczba rezerwacji, średni przychód
CREATE MATERIALIZED VIEW hotel_performance_summary AS
SELECT 
    h.id AS hotel_id,
    h.name AS hotel_name,
    h.star_standard,
    COALESCE(hr.current_rating, 0) AS current_rating,
    COALESCE(hr.review_count, 0) AS review_count,
    COUNT(DISTINCT r.id) AS total_reservations,
    COALESCE(AVG(r_rooms.base_price * (r.end_date - r.start_date)), 0) AS avg_revenue_per_reservation,
    COALESCE(SUM(r_rooms.base_price * (r.end_date - r.start_date)), 0) AS total_revenue
FROM hotels h
LEFT JOIN hotel_ratings hr ON h.id = hr.hotel_id
LEFT JOIN reservations r ON h.id = r.hotel_id
LEFT JOIN rooms r_rooms ON r.room_id = r_rooms.id
GROUP BY h.id, h.name, h.star_standard, hr.current_rating, hr.review_count;

-- Index dla lepszej wydajności
CREATE INDEX idx_hotel_performance_hotel_id ON hotel_performance_summary (hotel_id);

-- Widok zmaterializowany: monthly_revenue_summary
-- Miesięczne podsumowanie przychodów z podziałem na hotele
CREATE MATERIALIZED VIEW monthly_revenue_summary AS
SELECT 
    h.id AS hotel_id,
    h.name AS hotel_name,
    DATE_TRUNC('month', r.start_date) AS revenue_month,
    COUNT(r.id) AS monthly_reservations,
    SUM(rooms.base_price * (r.end_date - r.start_date)) AS monthly_base_revenue,
    SUM(COALESCE(addon_cost.total_addon_cost, 0)) AS monthly_addon_revenue,
    SUM(rooms.base_price * (r.end_date - r.start_date)) + SUM(COALESCE(addon_cost.total_addon_cost, 0)) AS monthly_total_revenue,
    AVG(rooms.base_price * (r.end_date - r.start_date)) AS avg_reservation_value
FROM hotels h
LEFT JOIN reservations r ON h.id = r.hotel_id
LEFT JOIN rooms ON r.room_id = rooms.id
LEFT JOIN (
    SELECT 
        rao.reservation_id,
        SUM(ra.price) AS total_addon_cost
    FROM reservation_add_ons rao
    JOIN reservation_addons ra ON rao.addon_id = ra.id
    GROUP BY rao.reservation_id
) addon_cost ON r.id = addon_cost.reservation_id
WHERE r.start_date IS NOT NULL
GROUP BY h.id, h.name, DATE_TRUNC('month', r.start_date)
ORDER BY revenue_month DESC, monthly_total_revenue DESC;

-- Index dla lepszej wydajności
CREATE INDEX idx_monthly_revenue_hotel_month ON monthly_revenue_summary (hotel_id, revenue_month);
CREATE INDEX idx_monthly_revenue_month ON monthly_revenue_summary (revenue_month);

-- Procedura odświeżania widoków zmaterializowanych
CREATE OR REPLACE PROCEDURE refresh_materialized_views()
LANGUAGE plpgsql
AS $$
BEGIN
    REFRESH MATERIALIZED VIEW hotel_performance_summary;
    REFRESH MATERIALIZED VIEW monthly_revenue_summary;
END;
$$;