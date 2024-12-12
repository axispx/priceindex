CREATE TABLE IF NOT EXISTS prices (
    timestamp timestamptz,
    source VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    price VARCHAR(255) NOT NULL,
    PRIMARY KEY (timestamp)
);