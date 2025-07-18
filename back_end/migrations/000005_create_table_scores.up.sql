CREATE TABLE scores(
    id serial primary key,
    user_id integer REFERENCES users,
    activity_id integer REFERENCES activities,
    text_id integer REFERENCES texts,
    points integer not null DEFAULT 0,
    duration integer DEFAULT 0,
    errors integer DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_scores_changetimestamp BEFORE UPDATE
    ON scores FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();
