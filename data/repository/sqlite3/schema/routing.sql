CREATE TABLE IF NOT EXISTS routing (
    job_category TEXT NOT NULL PRIMARY KEY,
    queue_name TEXT NOT NULL,
    UNIQUE (job_category, queue_name)
);