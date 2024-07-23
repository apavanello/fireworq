CREATE TABLE IF NOT EXISTS queue_throttle (
    name TEXT NOT NULL PRIMARY KEY,
    max_dispatches_per_second REAL NOT NULL,
    max_burst_size INTEGER NOT NULL
);