SELECT add_continuous_aggregate_policy('hourly_prices',
  start_offset => INTERVAL '1 day',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '1 hour');

SELECT add_continuous_aggregate_policy('daily_prices',
  start_offset => INTERVAL '7 day',
  end_offset => INTERVAL '1 day',
  schedule_interval => INTERVAL '1 day');
