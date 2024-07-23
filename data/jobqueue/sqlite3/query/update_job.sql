UPDATE "{{.JobQueue}}"
SET grabber_id = NULL,
    status = 'claimed',
    next_try = CAST((STRFTIME('%s', 'now') || SUBSTR(strftime('%f', 'now'), 4, 3)) AS INTEGER) + ?,
    retry_count = ?,
    fail_count = ?
WHERE job_id = ?;