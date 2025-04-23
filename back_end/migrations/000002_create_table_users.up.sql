CREATE TYPE user_type AS ENUM ('admin', 'regular', 'generic');

CREATE TABLE users(
    id serial primary key,
    user_type user_type not null DEFAULT 'regular',
    username varchar(60) not null,
    passwd_hash char(255) not null,
    name varchar(60),
    email varchar(60) not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_users_changetimestamp BEFORE UPDATE
    ON users FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();
