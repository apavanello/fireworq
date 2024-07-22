SELECT job_id FROM `{{.JobQueue}}`
WHERE status = ?
  AND next_try > ?
  AND next_try <= ?
  AND job_id <= ?
ORDER BY next_try DESC, job_id DESC LIMIT
