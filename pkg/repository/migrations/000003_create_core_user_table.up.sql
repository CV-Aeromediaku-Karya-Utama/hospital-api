BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS core_user
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

INSERT INTO core_user (id, name, username, password, sex, email, status)
VALUES ('61815430-6a92-44b5-81cd-0dfdfd5cab3e', 'admin', 'admin',
        '$2a$14$5QGizUm0t7oaGyg1.zi4VO7LlePwmvK27QWWOG1LN5ATxkbNB6ZYG', 'male', 'admin@admin.admin', 1);

COMMIT;