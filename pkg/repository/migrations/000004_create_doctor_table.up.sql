CREATE TABLE IF NOT EXISTS doctor
(
    id         uuid                NOT NULL DEFAULT uuid_generate_v4(),
    primary key (id)
);