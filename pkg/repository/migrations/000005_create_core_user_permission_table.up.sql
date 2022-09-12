BEGIN;
CREATE TABLE IF NOT EXISTS core_user_permission
(
    id            SERIAL NOT NULL,
    user_id       uuid   NOT NULL,
    permission_id SERIAL NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_core_user_permission_user_id
        FOREIGN KEY (user_id) REFERENCES core_user (id),
    CONSTRAINT fk_core_user_permission_permission_id
        FOREIGN KEY (permission_id) REFERENCES core_permission (id)
);
COMMIT;