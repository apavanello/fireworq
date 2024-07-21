UPDATE `{{.JobQueue}}`
SET status = 'grabbed', grabber_id = 1
WHERE job_id IN

