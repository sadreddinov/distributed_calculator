CREATE TABLE IF NOT EXISTS expression 
(
    id uuid Default gen_random_uuid(),
    expression varchar(255) not null,
    result varchar(255) Default '?',
    created_at timestamp with time zone not null,
    solved_at timestamp with time zone,
    work_state varchar(255),
    computing_resource_id uuid Default gen_random_uuid()
);

CREATE TABLE IF NOT EXISTS computing_resource
(
    id uuid not null,
    work_state varchar(255) not null,
    last_ping_at timestamp without time zone not  null
);

