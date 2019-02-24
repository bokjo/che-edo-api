CREATE TABLE IF NOT EXISTS jobs
(
id SERIAL,
name TEXT NOT NULL,
priority INT NOT NULL,
state TEXT NOT NULL,	
started TIMESTAMP NOT NULL,
completed	TIMESTAMP NOT NULL,
CONSTRAINT jobs_pkey PRIMARY KEY (id)
);