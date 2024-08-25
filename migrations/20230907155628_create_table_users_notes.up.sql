CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE notes (
    id bigserial not null primary key,
    user_id bigserial not null,
    note varchar not null
);
