CREATE TABLE IF NOT EXISTS Pinatas (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    color VARCHAR(255),
    shape VARCHAR(255),
    contents text[] NOT NULL,
    is_broken BOOLEAN NOT NULL,
    weight DECIMAL(10, 2),
    height DECIMAL(10, 2),
    width DECIMAL(10, 2),
    depth DECIMAL(10, 2),
    version integer NOT NULL DEFAULT 1
    );
