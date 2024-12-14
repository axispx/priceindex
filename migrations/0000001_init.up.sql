CREATE TABLE IF NOT EXISTS prices (
    timestamp timestamptz DEFAULT CURRENT_TIMESTAMP,
    source VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    price DECIMAL(30,30) NOT NULL,
    PRIMARY KEY (timestamp, address)
);
