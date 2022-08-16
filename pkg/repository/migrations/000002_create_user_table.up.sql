CREATE TABLE IF NOT EXISTS "user"
(
    id         SERIAL             NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    username   VARCHAR(100)        NOT NULL,
    password   VARCHAR(100)        NOT NULL,
    sex        VARCHAR(20)         NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    role_id    integer             not null default 1,
    created_at TIMESTAMP           NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP           NULL,
    primary key (id),
    foreign key (role_id) references role(id)
);