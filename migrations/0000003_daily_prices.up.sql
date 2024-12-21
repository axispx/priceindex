CREATE MATERIALIZED VIEW daily_prices
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 day', timestamp) as day,
  address,
  AVG(price) as avg_price
FROM prices
GROUP by day, address;
