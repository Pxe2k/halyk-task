INSERT INTO shows (title, start_time, total_seats) 
VALUES ('Zootopia 2', '2024-12-25 19:30:00', 100);

SET @show_id = LAST_INSERT_ID();

INSERT INTO seats (show_id, seat_number, is_reserved)
WITH RECURSIVE numbers AS (
    SELECT 1 AS n
    UNION ALL
    SELECT n + 1 FROM numbers WHERE n < 100
)
SELECT @show_id, n, FALSE FROM numbers;