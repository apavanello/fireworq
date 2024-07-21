CREATE TABLE IF NOT EXISTS "{{.JobQueue}}" (
    "job_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "next_try" INTEGER NOT NULL,
    "grabber_id" INTEGER,
    "status" TEXT CHECK( "status" IN ('claimed', 'grabbed') ) NOT NULL DEFAULT 'claimed',
    "created_at" INTEGER NOT NULL,
    "retry_count" INTEGER NOT NULL DEFAULT 0,
    "retry_delay" INTEGER NOT NULL DEFAULT 0,
    "fail_count" INTEGER NOT NULL DEFAULT 0,
    "category" TEXT NOT NULL,
    "url" BLOB,
    "payload" BLOB,
    "timeout" INTEGER,

    CONSTRAINT "chk_status" CHECK ( "status" IN ('claimed', 'grabbed') )
);

CREATE INDEX IF NOT EXISTS "grab" ON "{{.JobQueue}}" ("status", "next_try");