CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user"
(
    id         uuid                NOT NULL DEFAULT uuid_generate_v4(),
    name       VARCHAR(100)        NULL,
    username   VARCHAR(100)        NOT NULL,
    password   VARCHAR(100)        NOT NULL,
    sex        VARCHAR(20)         NULL,
    email      VARCHAR(100) UNIQUE NULL,
    status     INT                 NOT NULL DEFAULT 0,
    created_at TIMESTAMP           NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP           NULL,
    primary key (id)
);