-- Active: 1714057451806@@127.0.0.1@5432@restapi_test
CREATE TABLE IF NOT EXISTS commandsdb (
	id bigserial not null primary key,
	script TEXT NOT NULL UNIQUE,
	result TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_script ON commandsdb(script);
