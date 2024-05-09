CREATE TABLE IF NOT EXISTS commandsdb (
	id bigserial not null primary key,
	script TEXT NOT NULL UNIQUE,
	result TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_script ON commandsdb(script);

--  убрать !!!
--SELECT * FROM commandsdb; 
--DELETE FROM commandsdb WHERE id =55