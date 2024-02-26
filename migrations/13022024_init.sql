-- +goose Up
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

CREATE table security_questions (
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

INSERT INTO linux_questions (question, description)
VALUES
    ('linux question 1', 'test linux description 1'),
    ('linux question 2', 'test linux description 2'),
    ('linux question 3', 'test linux description 3'),
    ('linux question 4', 'test linux description 4');

INSERT INTO k8s_guestions (question, description)
VALUES
    ('k8s question 1', 'test k8s description 1'),
    ('k8s question 2', 'test k8s description 2'),
    ('k8s question 3', 'test k8s description 3'),
    ('k8s question 4', 'test k8s description 4');

INSERT INTO network_questions (question, description)
VALUES
    ('network question 1', 'network test description 1'),
    ('network question 2', 'network test description 2'),
    ('network question 3', 'network test description 3'),
    ('network question 4', 'network test description 4');

INSERT INTO security_questions (question, description)
VALUES
    (' security question 1', ' security test description 1'),
    (' security question 2', ' security test description 2'),
    (' security question 3', ' security test description 3'),
    (' security question 4', ' security test description 4');

INSERT INTO container_questions (question, description)
VALUES
    ('container question 1', 'container test description 1'),
    ('container question 2', 'container test description 2'),
    ('container question 3', 'container test description 3'),
    ('container question 4', 'container test description 4');

INSERT INTO developer_guestions (question, description)
VALUES
    ('developer question 1', 'developer test description 1'),
    ('developer question 2', 'developer test description 2'),
    ('developer question 3', 'developer test description 3'),
    ('developer question 4', 'developer test description 4');

-- +goose Down
--drop table survey;
--drop table users;
--drop type question_table;
drop table linux_questions;
drop table k8s_guestions;
drop table network_questions;
drop table security_questions;
drop table container_questions;
drop table developer_guestions;