DROP TABLE IF EXISTS files;
CREATE TABLE files (
id SERIAL NOT NULL,
created_at timestamp,
updated_at timestamp,
path text,
from text,
user_id integer,
signed_by integer
);
ALTER TABLE files OWNER TO sendto_server;
