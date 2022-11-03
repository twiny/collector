
-- Domains table
CREATE TABLE IF NOT EXISTS "details" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "website" TEXT,
    "url" TEXT,
    "page_title" TEXT,
    "html_file" TEXT,
    "first_visit" DATE,
    "last_visit" DATE
);

