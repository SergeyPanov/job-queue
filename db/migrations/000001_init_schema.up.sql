CREATE TYPE state AS ENUM ('succeeded', 'failed', 'progressing', 'unprocessed');

CREATE TABLE jobs
(
    id    SERIAL PRIMARY KEY,
    start_at timestamp not null not null default now() + interval '5' second,
    execution_time int,  -- interval is better
    state state NOT NULL DEFAULT 'unprocessed'::state,
    success_probability float default 0.5,
    attempts integer NOT NULL DEFAULT 0
);

