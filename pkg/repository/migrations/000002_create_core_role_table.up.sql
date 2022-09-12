BEGIN;
CREATE TABLE IF NOT EXISTS core_role
(
    id   SERIAL       NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);
INSERT INTO core_role (name)
VALUES ('ADMIN'),
       ('USER');
COMMIT;