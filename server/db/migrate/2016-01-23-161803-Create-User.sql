DROP TABLE IF EXISTS users;
CREATE TABLE users (
id SERIAL NOT NULL,
created_at timestamp,
updated_at timestamp,
password text,
status integer,
name text,
summary text,
key text,
email text,
role integer
);
ALTER TABLE users OWNER TO sendto_server;
