CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "tours" (
"id" TEXT PRIMARY KEY,
"title" TEXT NOT NULL,
"desc" TEXT NOT NULL,
"image_main" TEXT NOT NULL,
"content" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "featureds" (
"id" TEXT PRIMARY KEY,
"tour_id" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "groups" (
"id" TEXT PRIMARY KEY,
"title" TEXT NOT NULL,
"image_main" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "group_tours" (
"id" TEXT PRIMARY KEY,
"group_id" TEXT NOT NULL,
"tour_id" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
