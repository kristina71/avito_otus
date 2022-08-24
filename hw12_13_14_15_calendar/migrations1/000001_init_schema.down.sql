-- +migrate Down
drop index if exists event_owner_index;
drop index if exists event_times_index;
drop table if exists events;