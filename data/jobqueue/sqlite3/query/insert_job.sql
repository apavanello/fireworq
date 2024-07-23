INSERT INTO "{{.JobQueue}}" (next_try, created_at, retry_count, retry_delay, fail_count, category, url, payload, timeout)
VALUES (
           CAST((STRFTIME('%s', 'now') || SUBSTR(strftime('%f', 'now'), 4, 3)) AS INTEGER) + ?,
           CAST((STRFTIME('%s', 'now') || SUBSTR(strftime('%f', 'now'), 4, 3)) AS INTEGER),
           ?, ?, ?, ?, ?, ?, ?
)