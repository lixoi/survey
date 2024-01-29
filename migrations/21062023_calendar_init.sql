-- +goose Up
CREATE table surveys (
    id              serial primary key,
    user_id         number,
    title           text,
    question_table  text,
    question_id     text,
    answer          text,
    created_at      timestamptz not null default now(),
    latency         timestamptz,
    exist_to        timestamptz,
);

CREATE users_id (
    id                  serial primary key,
    prefer_questions    text as enum(),      
    surveys_done        boolean not null default 0,
    created_at          timestamptz not null default now(),
    exist_to            timestamptz
) 

CREATE linux_questions (
    id          serial primary key,
    question    text,
    description text
)

CREATE k8s_guestions (
    id          serial primary key,
    question    text,
    description text

CREATE network_questions (
    id          serial primary key,
    question    text,
    description text
)

CREATE information_security_questions (
    id          serial primary key,
    question    text,
    description text
)

CREATE container_questions (
    id          serial primary key,
    question    text,
    description text
)

CREATE developer_guestions (
    id          serial primary key,
    question    text,
    description text
) 

INSERT INTO events (id, title, created_at, description)
VALUES
    (1, 'test 1', now(), 'test description 1'),
    (2, 'test 2', now(), 'test description 2'),
    (3, 'test 3', now(), 'test description 3');

-- +goose Down
drop table events;
