-- Active: 1714057451806@@127.0.0.1@5432@restapi_script
CREATE TABLE IF NOT EXISTS commandsdb (
	id bigserial not null primary key,
	name TEXT NOT NULL UNIQUE,
	script TEXT NOT NULL,
	result TEXT NOT NULL);

--  убрать !!!
SELECT * FROM commandsdb; 