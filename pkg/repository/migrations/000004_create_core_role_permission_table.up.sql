BEGIN;
CREATE TABLE IF NOT EXISTS core_role_permission
(
    id            SERIAL NOT NULL,
    role_id       SERIAL NOT NULL,
    permission_id SERIAL NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_core_role_permission_role_id
        FOREIGN KEY (role_id) REFERENCES core_role (id),
    CONSTRAINT fk_core_role_permission_permission_id
        FOREIGN KEY (permission_id) REFERENCES core_permission (id)
);
COMMIT;