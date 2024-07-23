CREATE TABLE IF NOT EXISTS queue (
    name TEXT NOT NULL PRIMARY KEY,
    polling_interval INTEGER NOT NULL,
    max_workers INTEGER NOT NULL
);