CREATE TABLE seats (
    id SERIAL PRIMARY KEY,
    show_id INT NOT NULL REFERENCES shows(id) ON DELETE CASCADE,
    seat_number INT NOT NULL CHECK (seat_number > 0),
    is_reserved BOOLEAN NOT NULL DEFAULT false,
    reserved_at TIMESTAMP,
    UNIQUE (show_id, seat_number)
);