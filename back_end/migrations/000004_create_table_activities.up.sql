CREATE TABLE activities(
    id serial primary key,
    name varchar(60) not null,
    description text not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TRIGGER update_activities_changetimestamp BEFORE UPDATE
    ON activities FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();


INSERT INTO activities (name, description)
VALUES (
    'speed drill',
    'Type the text rows on screen as fast and accurately'
    'as you can until they stop. Your time to complete vs accuracy'
    'will determine your final score.'
)
