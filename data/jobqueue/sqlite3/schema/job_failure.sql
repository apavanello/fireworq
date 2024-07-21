CREATE TABLE IF NOT EXISTS "{{.Failure}}" (
    "failure_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "job_id" INTEGER NOT NULL,
    "category" TEXT NOT NULL,
    "url" BLOB,
    "payload" BLOB,
    "result" BLOB,
    "fail_count" INTEGER NOT NULL,
    "failed_at" INTEGER NOT NULL,
    "created_at" INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS "creation_order" ON "{{.Failure}}" ("created_at");