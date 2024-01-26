-- +goose Up
CREATE table events (
    id              serial primary key,
    title           text,
    created_at      timestamptz not null default now(),
    exist_to        timestamptz,
    description     text,
    user_id         text,
    time_send_report timestamptz
);

INSERT INTO events (id, title, created_at, description)
VALUES
    (1, 'test 1', now(), 'test description 1'),
    (2, 'test 2', now(), 'test description 2'),
    (3, 'test 3', now(), 'test description 3');

-- +goose Down
drop table events;
