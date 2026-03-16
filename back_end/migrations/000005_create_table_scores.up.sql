CREATE TABLE scores(
    id serial primary key,
    user_id integer REFERENCES users,
    activity_id integer REFERENCES activities,
    text_id integer REFERENCES texts,
    duration integer DEFAULT 0,
    result jsonb not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_scores_changetimestamp BEFORE UPDATE
    ON scores FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();
