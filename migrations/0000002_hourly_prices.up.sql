CREATE MATERIALIZED VIEW hourly_prices
WITH (timescaledb.continuous) AS
SELECT
  time_bucket('1 hour', timestamp) as hour,
  address,
  AVG(price) as avg_price
FROM prices
GROUP by hour, address;
