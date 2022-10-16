-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id UUID NOT NULL PRIMARY KEY,
    title varchar(256) not null,
    start_at timestamptz,
    duration bigint NOT NULL,
    description text,
    user_id UUID,
    remind_at bigint
);

CREATE INDEX if not exists event_times_index ON events (start_at, duration);

-- +goose Down
drop index if exists event_times_index;
drop table if exists events;
