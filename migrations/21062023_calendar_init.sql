-- +goose Up
CREATE TYPE question_table AS ENUM (
 'linux_questions', 
 'k8s_guestions', 
 'network_questions'
 'information_security_questions',
 'container_questions',
 'developer_guestions'
 );

CREATE table survey (
    id              serial primary key,
    user_id         number,
    title           text,
    question        text,
    answer          text,
    answered_at     timestamptz not null default now(),
    question_number number
);

CREATE users (
    id                        serial primary key,
    base_questions            question_table,      
    first_profile_questions   question_table,
    sec_profile_questions     question_table,
    survey_done               boolean not null default 0,
    created_at                timestamptz not null default now(),
    exist_to                  timestamptz not null default now()+interval '3 day'
); 

CREATE linux_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE k8s_guestions (
    id          serial primary key,
    question    text,
    description text
);

CREATE network_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE information_security_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE container_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE developer_guestions (
    id          serial primary key,
    question    text,
    description text
);

INSERT INTO events (id, title, created_at, description)
VALUES
    (1, 'test 1', now(), 'test description 1'),
    (2, 'test 2', now(), 'test description 2'),
    (3, 'test 3', now(), 'test description 3');

-- +goose Down
drop table surveys;
