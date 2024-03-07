-- +goose Down
 CREATE TYPE question_table AS ENUM (
 'LINUX_QUESTIONS', 
 'K8S_QUESTIONS', 
 'NETWORK_QUESTIONS',
 'SECURITY_QUESTIONS',
 'CONTAINER_QUESTIONS',
 'DEVELOPER_QUESTIONS'
 );

CREATE table survey (
    id              serial primary key,
    user_id         integer,
    title           text default '',
    question        text default '',
    answer          text default '',
    answered_at     timestamptz default null,
    question_number integer
);

CREATE table users (
    id                        serial primary key,
    base_questions            question_table,      
    first_profile_questions   question_table,
    sec_profile_questions     question_table,
    survey_done               boolean not null default false,
    created_at                timestamptz not null default now(),
    survey_start              timestamptz default null,  
    exist_to                  timestamptz not null default now()+interval '3 day'
); 
