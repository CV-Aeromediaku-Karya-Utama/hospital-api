CREATE TABLE IF NOT EXISTS weight
(
    id                   serial PRIMARY KEY,
    created_at           TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at           timestamp NULL,
    weight               int   not null,
    bmr                  int   not null,
    daily_caloric_intake int,
    user_id              int   not null,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);