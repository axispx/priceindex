CREATE TABLE IF NOT EXISTS market_caps (
    timestamp timestamptz DEFAULT CURRENT_TIMESTAMP,
    source VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    market_cap DECIMAL(60,30) NOT NULL,
    PRIMARY KEY (timestamp, address)
);

SELECT create_hypertable('market_caps', 'timestamp', migrate_data => TRUE);
