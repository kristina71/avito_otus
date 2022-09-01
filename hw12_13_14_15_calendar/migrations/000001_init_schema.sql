-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id serial primary key,
    title varchar(256) not null,
    start_at timestamptz,
    end_at timestamptz,
    description text,
    user_id int,
    remind_at timestamptz
    --created_at timestamp default now()
);

CREATE INDEX if not exists event_times_index ON events (start_at, end_at);
-- CREATE INDEX event_owner_index ON events (created_at);

-- +goose Down
drop index if exists event_owner_index;
drop index if exists event_times_index;
drop table if exists events;
