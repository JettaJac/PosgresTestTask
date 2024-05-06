-- Active: 1714057451806@@127.0.0.1@5432@restapi_script

-- TODO: script должен быть уникальным, нейм убрать
CREATE TABLE IF NOT EXISTS commandsdb (
	id bigserial not null primary key,
	name TEXT NOT NULL UNIQUE,
	script TEXT NOT NULL,
	result TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_script ON commandsdb(script);

--  убрать !!!
SELECT * FROM commandsdb; 
DELETE FROM commandsdb WHERE id =55