create table users (
    id uuid primary key,
    name text
);

create table address (user_id uuid, line1 text, state text);
