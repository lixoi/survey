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

CREATE table users (
    id                        serial primary key,
    base_questions            question_table,      
    first_profile_questions   question_table,
    sec_profile_questions     question_table,
    survey_done               boolean not null default 0,
    created_at                timestamptz not null default now(),
    exist_to                  timestamptz not null default now()+interval '3 day'
); 

CREATE table linux_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE table k8s_guestions (
    id          serial primary key,
    question    text,
    description text
);

CREATE table network_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE table information_security_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE table container_questions (
    id          serial primary key,
    question    text,
    description text
);

CREATE table developer_guestions (
    id          serial primary key,
    question    text,
    description text
);

INSERT INTO linux_questions (title, description)
VALUES
    ('linux question 1', 'test linux description 1'),
    ('linux question 2', 'test linux description 2'),
    ('linux question 3', 'test linux description 3');

INSERT INTO k8s_questions (title, description)
VALUES
    ('k8s question 1', 'test k8s description 1'),
    ('k8s question 2', 'test k8s description 2'),
    ('k8s question 3', 'test k8s description 3');

INSERT INTO network_questions (title, description)
VALUES
    ('network question 1', 'network test description 1'),
    ('network question 2', 'network test description 2'),
    ('network question 3', 'network test description 3');

INSERT INTO information_security_questions (title, description)
VALUES
    ('information security question 1', 'information security test description 1'),
    ('information security question 2', 'information security test description 2'),
    ('information security question 3', 'information security test description 3');

INSERT INTO container_questions (title, description)
VALUES
    ('container question 1', 'container test description 1'),
    ('container question 2', 'container test description 2'),
    ('container question 3', 'container test description 3');

INSERT INTO developer_guestions (title, description)
VALUES
    ('developer question 1', 'developer test description 1'),
    ('developer question 2', 'developer test description 2'),
    ('developer question 3', 'developer test description 3');

INSERT INTO users (title, description)
VALUES
    ('developer question 1', 'developer test description 1'),
    ('developer question 2', 'developer test description 2'),
    ('developer question 3', 'developer test description 3');


-- +goose Down
drop table survey;
