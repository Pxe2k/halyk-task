CREATE TABLE shows (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    total_seats INT NOT NULL CHECK (total_seats > 0)
);