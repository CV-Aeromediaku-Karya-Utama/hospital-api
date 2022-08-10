CREATE TABLE IF NOT EXISTS weight
(
    id                   serial PRIMARY KEY,
    created_at           TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at           timestamp NULL,
    weight               integer   not null,
    bmr                  integer   not null,
    daily_caloric_intake integer,
    user_id              integer   not null,
    FOREIGN KEY (user_id) REFERENCES user (id)
);