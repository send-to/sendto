DROP TABLE IF EXISTS files;
CREATE TABLE files (
id SERIAL NOT NULL,
created_at timestamp,
updated_at timestamp,
sender_id integer,
status integer,
path text,
user_id integer,
sender text
);
ALTER TABLE files OWNER TO sendto_server;
