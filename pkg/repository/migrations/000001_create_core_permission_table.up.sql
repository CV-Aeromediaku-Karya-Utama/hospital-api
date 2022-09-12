BEGIN;
CREATE TABLE IF NOT EXISTS core_permission
(
    id   SERIAL       NOT NULL,
    name VARCHAR(100) NOT NULL,
    primary key (id)
);
INSERT INTO core_permission (name)
VALUES ('manage-user');
COMMIT;