SELECT job_id FROM "{{.JobQueue}}"
WHERE status = 'claimed'
  AND next_try <= CAST((STRFTIME('%s', 'now') || SUBSTR(strftime('%f', 'now'), 4, 3)) AS INTEGER)
ORDER BY next_try ASC
LIMIT
