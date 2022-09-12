BEGIN;
CREATE TABLE IF NOT EXISTS core_user_role
(
    id      SERIAL NOT NULL,
    user_id uuid   NOT NULL,
    role_id SERIAL NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_core_user_role_user_id
        FOREIGN KEY (user_id) REFERENCES core_user (id),
    CONSTRAINT fk_core_user_role_role_id
        FOREIGN KEY (role_id) REFERENCES core_role (id)
);
COMMIT;