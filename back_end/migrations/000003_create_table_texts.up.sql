CREATE TYPE difficulty AS ENUM ('easy', 'normal', 'hard');

CREATE TYPE text_type AS ENUM ('drill', 'full-text', 'article');

CREATE TABLE texts(
    id serial primary key,
    text_type text_type not null DEFAULT 'full-text',
    title varchar(60) not null,
    difficulty difficulty not null DEFAULT 'easy',
    text_body text not null,
    text_length integer not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_texts_changetimestamp BEFORE UPDATE
    ON texts FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();
